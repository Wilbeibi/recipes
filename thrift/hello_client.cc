#include "gen-cpp/HelloSvc.h"
#include <thrift/transport/TSocket.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <boost/make_shared.hpp>
#include <iostream>
#include <string>

using namespace apache::thrift::transport;
using namespace apache::thrift::protocol;
using boost::make_shared;

int main() {
    auto trans_ep = make_shared<TSocket>("localhost", 9090);
    auto trans_buf = make_shared<TBufferedTransport>(trans_ep);
    auto proto = make_shared<TBinaryProtocol>(trans_buf);
    HelloSvcClient client(proto);
    try {
        trans_ep->open();
        std::string msg;
        client.hello_func(msg);
        std::cout << "[Client] received: " << msg << std::endl;
        trans_ep->close();
    } catch (...) {
        std::cout << "Client exception" << std::endl;
    }
}

