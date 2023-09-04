import numpy as np
from numpy import linalg as LA
from common import *
from sklearn.decomposition import PCA


@timer(cnt=count)
def create_random_vector(n: int) -> np.ndarray:
    """
    Создает вектор с рандомными элементами [0;100)
    :param n: int - shape.
    :return:
    - <class 'numpy.ndarray'> - рандомный вектор.
    """

    return np.random.rand(n) * 100


@timer(cnt=count)
def scale_vector(vec: np.ndarray, alpha: float) -> np.ndarray:
    """
    Масштабирует переданный вектор на alpha.
    :param vec: np.ndarray - вектор для масштабирования;
    :param alpha: float - параметр масштабирования.
    """

