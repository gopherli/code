# gee

1. net/htpp
    - 将路由和处理方法注册到映射表（http.ListenAndServe） 
    - 定义路由映射的处理方法（http.HandleFunc）
    - 解析请求的路径，查找路由映射表（ServeHttp）
            ||
            ||
            ||
            \/
2. gee
    - HTTP基础：定义路由、注册路由、解析路由
    - 上下文：
    - 模板
    - 工具集
