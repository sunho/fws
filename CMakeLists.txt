cmake_minimum_required(VERSION 3.10)

project (bot_registry)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

find_package(Boost 1.53.0 REQUIRED)
find_package(OpenSSL REQUIRED)

if(NOT Boost_FOUND) 
	message (FATAL_ERROR "Couldn't find Boost")
endif()

if(NOT OpenSSL_FOUND)
	message (FATAL_ERROR "Couldn't find OpenSSL")
endif()

set(INCLUDE_DIR "./include")

set(CPPRESTSDK_INCLUDE_DIR "./lib/cpprestsdk/Release/include")
set(HEADER_SEARCH_PATHS ${CPPRESTSDK_INCLUDE_DIR} ${OPENSSL_INCLUDE_DIR} ${Boost_INCLUDE_DIR} ${INCLUDE_DIR})
include_directories(${HEADER_SEARCH_PATHS})

set(CPPRESTSDK_LIBRARY "${PROJECT_SOURCE_DIR}/lib/cpprestsdk/build.release/Binaries/libcpprest.a")
set(LIBRARIES_SEARCH_PATHS ${CPPRESTSDK_LIBRARY} ${OPENSSL_LIBRARIES} ${Boost_LIBRARIES})

add_executable(${PROJECT_NAME} ./src/main.cpp)

target_link_libraries(${PROJECT_NAME} ${LIBRARIES_SEARCH_PATHS})
