# Flatten embedded list
from compiler.ast import flatten
flatten([1, [2], [3, [[4]]])

# Get docstring from a function
def my_func():
    "Docstring blabla"
    pass
my_func.__doc__

# Get function signature
# https://docs.python.org/3.4/library/inspect.html#inspect.signature

# Call function by it's literal name
# https://stackoverflow.com/questions/3061/calling-a-function-of-a-module-from-a-string-with-the-functions-name-in-python

# Get list of methods in a class
# https://stackoverflow.com/questions/1911281/how-do-i-get-list-of-methods-in-a-python-class
# inspect.getmembers()

# Get CPU numbers
from multiprocessing import cpu_count
print (cpu_count())

# From github.com/dongweiming/web_develop.git
def humanize_bytes(bytesize, precision=2):
    abbrevs = (
        (1 << 50, 'PB'),
        (1 << 40, 'TB'),
        (1 << 30, 'GB'),
        (1 << 20, 'MB'),
        (1 << 10, 'kB'),
        (1, 'bytes')
    )
    if bytesize == 1:
        return '1 byte'
    for factor, suffix in abbrevs:
        if bytesize >= factor:
            break
    return '%.*f %s' % (precision, bytesize / factor, suffix)

# progress bar
from tqdm import tqdm
urls = ['url1', 'url2', 'url3']
for url in tgdm(urls):
    # process(url)