import time

def excution_time(f):
    def decorator(*args, **kwargs):
        begin = time.time()
        ret = f(*args, **kwargs)

        print(f"time take in function: {f.__name__} = {time.time() - begin}")
        return ret

    return decorator