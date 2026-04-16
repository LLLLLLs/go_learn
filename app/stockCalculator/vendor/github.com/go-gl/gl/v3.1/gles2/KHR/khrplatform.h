#ifndef __khrplatform_h_
#define __khrplatform_h_

#include <stddef.h>
#include <stdint.h>

#define KHRONOS_APICALL
#define KHRONOS_APIENTRY
#define KHRONOS_APIATTRIBUTES

#define KHRONOS_SUPPORT_INT64 1
#define KHRONOS_SUPPORT_FLOAT 1

typedef int8_t khronos_int8_t;
typedef uint8_t khronos_uint8_t;
typedef int16_t khronos_int16_t;
typedef uint16_t khronos_uint16_t;
typedef int32_t khronos_int32_t;
typedef uint32_t khronos_uint32_t;
typedef int64_t khronos_int64_t;
typedef uint64_t khronos_uint64_t;
typedef intptr_t khronos_intptr_t;
typedef uintptr_t khronos_uintptr_t;
typedef ptrdiff_t khronos_ssize_t;
typedef size_t khronos_usize_t;
typedef float khronos_float_t;
typedef uint64_t khronos_utime_nanoseconds_t;
typedef int64_t khronos_stime_nanoseconds_t;
typedef uint64_t khronos_time_ns_t;

typedef enum {
	KHRONOS_FALSE = 0,
	KHRONOS_TRUE = 1,
} khronos_boolean_enum_t;

#endif
