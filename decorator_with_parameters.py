permitted_role = ['admin']


def check_access(role):
    def inner(func):
        def wrapper():
            #print('decorator on')
            try:
                if role in permitted_role:
                    return func()
                else:
                    raise PermissionError
            except PermissionError:
                return f"PermissionError: Access denied for user {role}"
            #print('decorator off')
        return wrapper
    return inner


@check_access('admin')
def user_login():
    return "Добро пожаловать!"


print(user_login())


# Результат работы: Добро пожаловать!


@check_access('user')
def user_login():
    return "Добро пожаловать!"


print(user_login())
# Результат работы: PermissionError: Access denied for user "user"
