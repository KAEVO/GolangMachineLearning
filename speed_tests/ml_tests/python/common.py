# Определение декораторов
# и данных для тестирования

import os
import time
from memory_profiler import profile

DEFAULT = 10

try:
    count = os.environ["PYTHON_TESTS_COUNT"]
except KeyError:
    count = DEFAULT


def timer(cnt):
    """
    Таймер для замера скорости работы тестируемых функций.
    :param cnt: int - кол-во прогонов функции.
    """

    def wrp(func):
        def wrapper(*args, **kwargs):
            start = time.time()
            print(f"Start: {start}")
            for n in range(cnt):
  