# *** Project

## introduce

- Integration of pprof, cors, recovery, access_log, gzip and other extensions of Hertz.
- Generating the base code for unit tests.
- Provides basic profile functions.
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| biz/hertz_gen  | Hertz generated model |
| biz/handler  | Used for request processing, validation and return of response. |
| biz/service  | The actual business logic. |
| biz/dal  | Logic for operating the storage layer |
| biz/route  | Routing and middleware registration |