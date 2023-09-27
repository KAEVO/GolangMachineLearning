#include <iostream>
#include <mlpack_tests.hpp>

using namespace std;

void linear_fit(int rep){
    arma::mat x_train;
    arma::vec y_train;

    data::Load("../../../datasets/the_WWT_weather_dataset_xtrain.csv", x_train, true);
    y_train.load("../../../datasets/the_WWT_weather_dataset_ytrain.csv");

    LinearRegression lr;
    double diff = 0.0;
    for (int i = 0; i < rep; i++){
        auto start = chrono::steady_clock::now();
        lr.Train(x_train, y_train);
        auto end = chrono::steady_clock::now();
        diff += chrono::duration <double, nano> (end - start).count();
    }
    cout << fixed << "Test duration : " << diff/(float)rep << " ns" << endl;
}

void logit_fit(int rep){
    arma::mat x_train;
    arma::mat y_train;

    data::Load("../../../datasets/the_breast_canser_dataset_xtrain.csv", x_train, true);
    data::Load("../../../datasets/the_breast_canser_dataset_ytrain.csv", y_train, true);

    arma::Row<size_t> responses(y_train.n_cols);

    for (int i = 0; i < y_train.n_cols; i++){
        responses[i] = (size_t)y_train.row(0)[i];
    }

    LogisticRegression<> logit(x_train.n_rows);
    doub