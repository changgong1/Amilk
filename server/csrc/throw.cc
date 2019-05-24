#include <iostream>
using namespace std;
void Throw(){throw 1;}
void NoBlockThrow(){Throw();}
void BlockThrow() throw() { Throw();};

int main(){
    try{
    Throw();
    }
    catch(...){
        cout << "found throw." << endl;
    }
    try{
    NoBlockThrow();
    }
    catch(...){
        cout << "-----------------" << endl;
        cout << "throw is not blocked." << endl;
    }
    try{
    BlockThrow();//terminate called after throwing an instance of'int'
    }
    catch(...){
        cout << "found throw 1." << endl;
    }
}