// hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod1(1: HelloReq request) (api.get="/hello1");
    HelloResp HelloMethod2(1: HelloReq request) (api.get="/hello2");
    HelloResp HelloMethod3(1: HelloReq request) (api.get="/hello3");
}
