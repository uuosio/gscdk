
add_cython_target(hello CXX)

add_library(_gscdk MODULE
    ${hello}
)

python_extension_module(_gscdk)

target_include_directories(_gscdk PRIVATE 
    ${CMAKE_CURRENT_SOURCE_DIR}/uuoskit
)

install(TARGETS _gscdk LIBRARY DESTINATION pysrc)

