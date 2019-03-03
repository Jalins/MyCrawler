# MyCrawler
这里的源码都是根据教学视频写出来的，只供学习使用。

## 单任务版
### 1.总体思路
![1551600468944](E:\GoWorkspace\src\MyCrawler\images\单任务版爬虫思路.png)

### 2.具体思路

**1）**上面是总体的一个思路，这里是具体到代码层面如何实现。首先是数据结构需要满足一个可嵌套的特性，再者就是这个数据结构需要满足一个需求：输入与输出是可同时进行。基于这两个需求，给出下面的结构设计：

```go
// 一个requests包含一个Url以及一个函数，这样就能保证在程序中一个输入对应着一个输出。
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}
// 这个设计可以达到嵌套的目的，因为程序的需求是一个页面中包含了另一个页面，这个页面又包含着另外一个页面，至于为什么是切片的形式，这是因为一个页面中可能包含多页面。
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func ParserFunc([]byte) ParseResult {
	return ParseResult{}
}
```

**2）**结构体设计后接下来就是程序的设计，程序大概可以分为四个部分，分别为engine、队列、parse以及fetch。

- **fetch函数**，fetch的功能很简单就是根据传入的url参数通过Get/Post方法从互联网上爬取到数据。大概的代码如下：

  ```go
  func Fetch(url string) ([]byte, error) {
  	resp, err := http.Get(url)
  	if err != nil {
  		return nil, err
  	}
  	defer resp.Body.Close()
  
  	if resp.StatusCode != http.StatusOK {
  		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
  	}
  	readall, err := ioutil.ReadAll(resp.Body)
  
  	return readall, nil
  
  }
  ```

- **parse函数**，parse函数的作用便是将fetch函数返回的结果进行解析，使用正则表达式从大量的数据中获取到有用的数据，将这些数据封装到ParseResult{}中并返回出去。代码如下：

  ```go
  var citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]*)</a>`
  
  func ParseCityList(contents []byte) engine.ParseResult {
  
  	re := regexp.MustCompile(citylistRe)
  
  	matches := re.FindAllSubmatch(contents, -1)
  	result := engine.ParseResult{}
  
  	for _, value := range matches {
  		result.Items = append(result.Items, "City "+string(value[2]))
  		result.Requests = append(result.Requests, engine.Request{
  			Url:        string(value[1]),
  			ParserFunc: ParseCity,
  		})
  	}
  	return result
  }
  ```

- **队列**，队列的目的是保存一个个Request，供给engine进行分析，关于的队列的代码很简单，如下：

  ```go
  // 定义一个切片（队列），类型为Request
  var requests []Request
  
  // 不断的取出第一个元素
  r := requests[0]        
  requests = requests[1:]
  
  // 将一个个Request添加到队列中去
  requests = append(requests, parseResult.Requests...)
  ```

- **engine**，引擎在程序中起到了中枢的作用，将队列，parse以及fetch联系在一起。代码逻辑：

  将传入的Request添加进一个切片队列中，然后从这个队列中获取到第一个元素（即刚才传入的Request）；**1**.接着先根据这个Request的Url属性使用fetch函数从网上获取到数据text；**2.**然后将得到的text传递给parse函数，parse函数开始从text中的大量信息提取到有用的数据（比如从一个城市列表中提取到一个个城市item以及对应的url，其他的信息丢弃），将这些信息组合成一个ParseResult，这个ParseResult包含了一个个item以及一个个Request{Url, Parser}；**3.**将ParseResult中的所有Request都添加进队列中，engine不断的从这个队列中去元素进行分析。


