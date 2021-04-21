#include <fstream>
#include <iostream>
#include <json/json.h>
#include <cassert>
#include <errno.h>
#include <string.h>
#include "opencv2/opencv.hpp"

using namespace cv;
using namespace std;
using namespace Json;
//Jsoncpp 测试
//int main(void)
//{
//	CharReaderBuilder builder;
//	builder["collectComments"] = false;
//	Value sin;
//	string errs;
//	Value root;
//	ifstream ifs;
//	ifs.open("test.json");
//	parseFromStream(builder, ifs, &root, &errs);
//	cout << root << endl;
//	cout << root["asd"] << endl;
//	if (root["names"] == nullValue) {
//		cout << "123" << endl;
//	}
//	ifs.close();
//	return 0;
//}

//opencv测试
int main() {
	Mat image;
	string fileName = "C:/Users/MSI_NB/Desktop/hardware.png";
	image = imread(fileName);
	namedWindow("test", WINDOW_AUTOSIZE);
	imshow("test", image);
	waitKey(0);
	destroyWindow("test");

	return 0;
}
