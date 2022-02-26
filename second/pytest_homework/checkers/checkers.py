def smoke_check(func):
    def wrapper(*args, **kwargs):
        print('decorator on')
        func(*args, **kwargs)
        print('decorator off')
    return wrapper

