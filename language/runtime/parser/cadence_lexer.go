// Code generated from parser/Cadence.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 73, 531,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52,
	3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3,
	54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60,
	3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 62, 3, 62, 7, 62, 411, 10, 62, 12, 62, 14, 62, 414,
	11, 62, 3, 63, 5, 63, 417, 10, 63, 3, 64, 3, 64, 5, 64, 421, 10, 64, 3,
	65, 3, 65, 7, 65, 425, 10, 65, 12, 65, 14, 65, 428, 11, 65, 3, 66, 3, 66,
	3, 66, 3, 66, 6, 66, 434, 10, 66, 13, 66, 14, 66, 435, 3, 67, 3, 67, 3,
	67, 3, 67, 6, 67, 442, 10, 67, 13, 67, 14, 67, 443, 3, 68, 3, 68, 3, 68,
	3, 68, 6, 68, 450, 10, 68, 13, 68, 14, 68, 451, 3, 69, 3, 69, 3, 69, 7,
	69, 457, 10, 69, 12, 69, 14, 69, 460, 11, 69, 3, 70, 3, 70, 7, 70, 464,
	10, 70, 12, 70, 14, 70, 467, 11, 70, 3, 70, 3, 70, 3, 71, 3, 71, 5, 71,
	473, 10, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 6, 72, 482,
	10, 72, 13, 72, 14, 72, 483, 3, 72, 3, 72, 5, 72, 488, 10, 72, 3, 73, 3,
	73, 3, 74, 6, 74, 493, 10, 74, 13, 74, 14, 74, 494, 3, 74, 3, 74, 3, 75,
	6, 75, 500, 10, 75, 13, 75, 14, 75, 501, 3, 75, 3, 75, 3, 76, 3, 76, 3,
	76, 3, 76, 3, 76, 7, 76, 511, 10, 76, 12, 76, 14, 76, 514, 11, 76, 3, 76,
	3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77, 3, 77, 7, 77, 525, 10,
	77, 12, 77, 14, 77, 528, 11, 77, 3, 77, 3, 77, 3, 512, 2, 78, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 2, 127, 2, 129, 64,
	131, 65, 133, 66, 135, 67, 137, 68, 139, 69, 141, 2, 143, 2, 145, 2, 147,
	70, 149, 71, 151, 72, 153, 73, 3, 2, 15, 5, 2, 67, 92, 97, 97, 99, 124,
	3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 4, 2, 50, 51, 97, 97, 4, 2, 50, 57,
	97, 97, 6, 2, 50, 59, 67, 72, 97, 97, 99, 104, 4, 2, 67, 92, 99, 124, 6,
	2, 50, 59, 67, 92, 97, 97, 99, 124, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94,
	9, 2, 36, 36, 41, 41, 50, 50, 94, 94, 112, 112, 116, 116, 118, 118, 5,
	2, 50, 59, 67, 72, 99, 104, 6, 2, 2, 2, 11, 11, 13, 14, 34, 34, 4, 2, 12,
	12, 15, 15, 2, 541, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2,
	2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3,
	2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77,
	3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2,
	85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2,
	2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2,
	2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107,
	3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2,
	2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3,
	2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2,
	133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2,
	2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153,
	3, 2, 2, 2, 3, 155, 3, 2, 2, 2, 5, 157, 3, 2, 2, 2, 7, 159, 3, 2, 2, 2,
	9, 161, 3, 2, 2, 2, 11, 163, 3, 2, 2, 2, 13, 165, 3, 2, 2, 2, 15, 167,
	3, 2, 2, 2, 17, 169, 3, 2, 2, 2, 19, 173, 3, 2, 2, 2, 21, 175, 3, 2, 2,
	2, 23, 178, 3, 2, 2, 2, 25, 181, 3, 2, 2, 2, 27, 183, 3, 2, 2, 2, 29, 186,
	3, 2, 2, 2, 31, 189, 3, 2, 2, 2, 33, 191, 3, 2, 2, 2, 35, 193, 3, 2, 2,
	2, 37, 196, 3, 2, 2, 2, 39, 199, 3, 2, 2, 2, 41, 201, 3, 2, 2, 2, 43, 203,
	3, 2, 2, 2, 45, 205, 3, 2, 2, 2, 47, 207, 3, 2, 2, 2, 49, 209, 3, 2, 2,
	2, 51, 211, 3, 2, 2, 2, 53, 213, 3, 2, 2, 2, 55, 216, 3, 2, 2, 2, 57, 218,
	3, 2, 2, 2, 59, 222, 3, 2, 2, 2, 61, 225, 3, 2, 2, 2, 63, 229, 3, 2, 2,
	2, 65, 231, 3, 2, 2, 2, 67, 233, 3, 2, 2, 2, 69, 245, 3, 2, 2, 2, 71, 252,
	3, 2, 2, 2, 73, 261, 3, 2, 2, 2, 75, 270, 3, 2, 2, 2, 77, 280, 3, 2, 2,
	2, 79, 284, 3, 2, 2, 2, 81, 290, 3, 2, 2, 2, 83, 295, 3, 2, 2, 2, 85, 299,
	3, 2, 2, 2, 87, 304, 3, 2, 2, 2, 89, 309, 3, 2, 2, 2, 91, 313, 3, 2, 2,
	2, 93, 322, 3, 2, 2, 2, 95, 329, 3, 2, 2, 2, 97, 335, 3, 2, 2, 2, 99, 344,
	3, 2, 2, 2, 101, 348, 3, 2, 2, 2, 103, 352, 3, 2, 2, 2, 105, 355, 3, 2,
	2, 2, 107, 360, 3, 2, 2, 2, 109, 366, 3, 2, 2, 2, 111, 371, 3, 2, 2, 2,
	113, 377, 3, 2, 2, 2, 115, 381, 3, 2, 2, 2, 117, 388, 3, 2, 2, 2, 119,
	393, 3, 2, 2, 2, 121, 400, 3, 2, 2, 2, 123, 408, 3, 2, 2, 2, 125, 416,
	3, 2, 2, 2, 127, 420, 3, 2, 2, 2, 129, 422, 3, 2, 2, 2, 131, 429, 3, 2,
	2, 2, 133, 437, 3, 2, 2, 2, 135, 445, 3, 2, 2, 2, 137, 453, 3, 2, 2, 2,
	139, 461, 3, 2, 2, 2, 141, 472, 3, 2, 2, 2, 143, 487, 3, 2, 2, 2, 145,
	489, 3, 2, 2, 2, 147, 492, 3, 2, 2, 2, 149, 499, 3, 2, 2, 2, 151, 505,
	3, 2, 2, 2, 153, 520, 3, 2, 2, 2, 155, 156, 7, 61, 2, 2, 156, 4, 3, 2,
	2, 2, 157, 158, 7, 46, 2, 2, 158, 6, 3, 2, 2, 2, 159, 160, 7, 125, 2, 2,
	160, 8, 3, 2, 2, 2, 161, 162, 7, 127, 2, 2, 162, 10, 3, 2, 2, 2, 163, 164,
	7, 60, 2, 2, 164, 12, 3, 2, 2, 2, 165, 166, 7, 93, 2, 2, 166, 14, 3, 2,
	2, 2, 167, 168, 7, 95, 2, 2, 168, 16, 3, 2, 2, 2, 169, 170, 7, 62, 2, 2,
	170, 171, 7, 47, 2, 2, 171, 172, 7, 64, 2, 2, 172, 18, 3, 2, 2, 2, 173,
	174, 7, 63, 2, 2, 174, 20, 3, 2, 2, 2, 175, 176, 7, 126, 2, 2, 176, 177,
	7, 126, 2, 2, 177, 22, 3, 2, 2, 2, 178, 179, 7, 40, 2, 2, 179, 180, 7,
	40, 2, 2, 180, 24, 3, 2, 2, 2, 181, 182, 7, 48, 2, 2, 182, 26, 3, 2, 2,
	2, 183, 184, 7, 63, 2, 2, 184, 185, 7, 63, 2, 2, 185, 28, 3, 2, 2, 2, 186,
	187, 7, 35, 2, 2, 187, 188, 7, 63, 2, 2, 188, 30, 3, 2, 2, 2, 189, 190,
	7, 62, 2, 2, 190, 32, 3, 2, 2, 2, 191, 192, 7, 64, 2, 2, 192, 34, 3, 2,
	2, 2, 193, 194, 7, 62, 2, 2, 194, 195, 7, 63, 2, 2, 195, 36, 3, 2, 2, 2,
	196, 197, 7, 64, 2, 2, 197, 198, 7, 63, 2, 2, 198, 38, 3, 2, 2, 2, 199,
	200, 7, 45, 2, 2, 200, 40, 3, 2, 2, 2, 201, 202, 7, 47, 2, 2, 202, 42,
	3, 2, 2, 2, 203, 204, 7, 44, 2, 2, 204, 44, 3, 2, 2, 2, 205, 206, 7, 49,
	2, 2, 206, 46, 3, 2, 2, 2, 207, 208, 7, 39, 2, 2, 208, 48, 3, 2, 2, 2,
	209, 210, 7, 40, 2, 2, 210, 50, 3, 2, 2, 2, 211, 212, 7, 35, 2, 2, 212,
	52, 3, 2, 2, 2, 213, 214, 7, 62, 2, 2, 214, 215, 7, 47, 2, 2, 215, 54,
	3, 2, 2, 2, 216, 217, 7, 65, 2, 2, 217, 56, 3, 2, 2, 2, 218, 219, 5, 147,
	74, 2, 219, 220, 7, 65, 2, 2, 220, 221, 7, 65, 2, 2, 221, 58, 3, 2, 2,
	2, 222, 223, 7, 99, 2, 2, 223, 224, 7, 117, 2, 2, 224, 60, 3, 2, 2, 2,
	225, 226, 7, 99, 2, 2, 226, 227, 7, 117, 2, 2, 227, 228, 7, 65, 2, 2, 228,
	62, 3, 2, 2, 2, 229, 230, 7, 42, 2, 2, 230, 64, 3, 2, 2, 2, 231, 232, 7,
	43, 2, 2, 232, 66, 3, 2, 2, 2, 233, 234, 7, 118, 2, 2, 234, 235, 7, 116,
	2, 2, 235, 236, 7, 99, 2, 2, 236, 237, 7, 112, 2, 2, 237, 238, 7, 117,
	2, 2, 238, 239, 7, 99, 2, 2, 239, 240, 7, 101, 2, 2, 240, 241, 7, 118,
	2, 2, 241, 242, 7, 107, 2, 2, 242, 243, 7, 113, 2, 2, 243, 244, 7, 112,
	2, 2, 244, 68, 3, 2, 2, 2, 245, 246, 7, 117, 2, 2, 246, 247, 7, 118, 2,
	2, 247, 248, 7, 116, 2, 2, 248, 249, 7, 119, 2, 2, 249, 250, 7, 101, 2,
	2, 250, 251, 7, 118, 2, 2, 251, 70, 3, 2, 2, 2, 252, 253, 7, 116, 2, 2,
	253, 254, 7, 103, 2, 2, 254, 255, 7, 117, 2, 2, 255, 256, 7, 113, 2, 2,
	256, 257, 7, 119, 2, 2, 257, 258, 7, 116, 2, 2, 258, 259, 7, 101, 2, 2,
	259, 260, 7, 103, 2, 2, 260, 72, 3, 2, 2, 2, 261, 262, 7, 101, 2, 2, 262,
	263, 7, 113, 2, 2, 263, 264, 7, 112, 2, 2, 264, 265, 7, 118, 2, 2, 265,
	266, 7, 116, 2, 2, 266, 267, 7, 99, 2, 2, 267, 268, 7, 101, 2, 2, 268,
	269, 7, 118, 2, 2, 269, 74, 3, 2, 2, 2, 270, 271, 7, 107, 2, 2, 271, 272,
	7, 112, 2, 2, 272, 273, 7, 118, 2, 2, 273, 274, 7, 103, 2, 2, 274, 275,
	7, 116, 2, 2, 275, 276, 7, 104, 2, 2, 276, 277, 7, 99, 2, 2, 277, 278,
	7, 101, 2, 2, 278, 279, 7, 103, 2, 2, 279, 76, 3, 2, 2, 2, 280, 281, 7,
	104, 2, 2, 281, 282, 7, 119, 2, 2, 282, 283, 7, 112, 2, 2, 283, 78, 3,
	2, 2, 2, 284, 285, 7, 103, 2, 2, 285, 286, 7, 120, 2, 2, 286, 287, 7, 103,
	2, 2, 287, 288, 7, 112, 2, 2, 288, 289, 7, 118, 2, 2, 289, 80, 3, 2, 2,
	2, 290, 291, 7, 103, 2, 2, 291, 292, 7, 111, 2, 2, 292, 293, 7, 107, 2,
	2, 293, 294, 7, 118, 2, 2, 294, 82, 3, 2, 2, 2, 295, 296, 7, 114, 2, 2,
	296, 297, 7, 116, 2, 2, 297, 298, 7, 103, 2, 2, 298, 84, 3, 2, 2, 2, 299,
	300, 7, 114, 2, 2, 300, 301, 7, 113, 2, 2, 301, 302, 7, 117, 2, 2, 302,
	303, 7, 118, 2, 2, 303, 86, 3, 2, 2, 2, 304, 305, 7, 114, 2, 2, 305, 306,
	7, 116, 2, 2, 306, 307, 7, 107, 2, 2, 307, 308, 7, 120, 2, 2, 308, 88,
	3, 2, 2, 2, 309, 310, 7, 114, 2, 2, 310, 311, 7, 119, 2, 2, 311, 312, 7,
	100, 2, 2, 312, 90, 3, 2, 2, 2, 313, 314, 7, 114, 2, 2, 314, 315, 7, 119,
	2, 2, 315, 316, 7, 100, 2, 2, 316, 317, 7, 42, 2, 2, 317, 318, 7, 117,
	2, 2, 318, 319, 7, 103, 2, 2, 319, 320, 7, 118, 2, 2, 320, 321, 7, 43,
	2, 2, 321, 92, 3, 2, 2, 2, 322, 323, 7, 116, 2, 2, 323, 324, 7, 103, 2,
	2, 324, 325, 7, 118, 2, 2, 325, 326, 7, 119, 2, 2, 326, 327, 7, 116, 2,
	2, 327, 328, 7, 112, 2, 2, 328, 94, 3, 2, 2, 2, 329, 330, 7, 100, 2, 2,
	330, 331, 7, 116, 2, 2, 331, 332, 7, 103, 2, 2, 332, 333, 7, 99, 2, 2,
	333, 334, 7, 109, 2, 2, 334, 96, 3, 2, 2, 2, 335, 336, 7, 101, 2, 2, 336,
	337, 7, 113, 2, 2, 337, 338, 7, 112, 2, 2, 338, 339, 7, 118, 2, 2, 339,
	340, 7, 107, 2, 2, 340, 341, 7, 112, 2, 2, 341, 342, 7, 119, 2, 2, 342,
	343, 7, 103, 2, 2, 343, 98, 3, 2, 2, 2, 344, 345, 7, 110, 2, 2, 345, 346,
	7, 103, 2, 2, 346, 347, 7, 118, 2, 2, 347, 100, 3, 2, 2, 2, 348, 349, 7,
	120, 2, 2, 349, 350, 7, 99, 2, 2, 350, 351, 7, 116, 2, 2, 351, 102, 3,
	2, 2, 2, 352, 353, 7, 107, 2, 2, 353, 354, 7, 104, 2, 2, 354, 104, 3, 2,
	2, 2, 355, 356, 7, 103, 2, 2, 356, 357, 7, 110, 2, 2, 357, 358, 7, 117,
	2, 2, 358, 359, 7, 103, 2, 2, 359, 106, 3, 2, 2, 2, 360, 361, 7, 121, 2,
	2, 361, 362, 7, 106, 2, 2, 362, 363, 7, 107, 2, 2, 363, 364, 7, 110, 2,
	2, 364, 365, 7, 103, 2, 2, 365, 108, 3, 2, 2, 2, 366, 367, 7, 118, 2, 2,
	367, 368, 7, 116, 2, 2, 368, 369, 7, 119, 2, 2, 369, 370, 7, 103, 2, 2,
	370, 110, 3, 2, 2, 2, 371, 372, 7, 104, 2, 2, 372, 373, 7, 99, 2, 2, 373,
	374, 7, 110, 2, 2, 374, 375, 7, 117, 2, 2, 375, 376, 7, 103, 2, 2, 376,
	112, 3, 2, 2, 2, 377, 378, 7, 112, 2, 2, 378, 379, 7, 107, 2, 2, 379, 380,
	7, 110, 2, 2, 380, 114, 3, 2, 2, 2, 381, 382, 7, 107, 2, 2, 382, 383, 7,
	111, 2, 2, 383, 384, 7, 114, 2, 2, 384, 385, 7, 113, 2, 2, 385, 386, 7,
	116, 2, 2, 386, 387, 7, 118, 2, 2, 387, 116, 3, 2, 2, 2, 388, 389, 7, 104,
	2, 2, 389, 390, 7, 116, 2, 2, 390, 391, 7, 113, 2, 2, 391, 392, 7, 111,
	2, 2, 392, 118, 3, 2, 2, 2, 393, 394, 7, 101, 2, 2, 394, 395, 7, 116, 2,
	2, 395, 396, 7, 103, 2, 2, 396, 397, 7, 99, 2, 2, 397, 398, 7, 118, 2,
	2, 398, 399, 7, 103, 2, 2, 399, 120, 3, 2, 2, 2, 400, 401, 7, 102, 2, 2,
	401, 402, 7, 103, 2, 2, 402, 403, 7, 117, 2, 2, 403, 404, 7, 118, 2, 2,
	404, 405, 7, 116, 2, 2, 405, 406, 7, 113, 2, 2, 406, 407, 7, 123, 2, 2,
	407, 122, 3, 2, 2, 2, 408, 412, 5, 125, 63, 2, 409, 411, 5, 127, 64, 2,
	410, 409, 3, 2, 2, 2, 411, 414, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 412,
	413, 3, 2, 2, 2, 413, 124, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 415, 417,
	9, 2, 2, 2, 416, 415, 3, 2, 2, 2, 417, 126, 3, 2, 2, 2, 418, 421, 9, 3,
	2, 2, 419, 421, 5, 125, 63, 2, 420, 418, 3, 2, 2, 2, 420, 419, 3, 2, 2,
	2, 421, 128, 3, 2, 2, 2, 422, 426, 9, 3, 2, 2, 423, 425, 9, 4, 2, 2, 424,
	423, 3, 2, 2, 2, 425, 428, 3, 2, 2, 2, 426, 424, 3, 2, 2, 2, 426, 427,
	3, 2, 2, 2, 427, 130, 3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 429, 430, 7, 50,
	2, 2, 430, 431, 7, 100, 2, 2, 431, 433, 3, 2, 2, 2, 432, 434, 9, 5, 2,
	2, 433, 432, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 435,
	436, 3, 2, 2, 2, 436, 132, 3, 2, 2, 2, 437, 438, 7, 50, 2, 2, 438, 439,
	7, 113, 2, 2, 439, 441, 3, 2, 2, 2, 440, 442, 9, 6, 2, 2, 441, 440, 3,
	2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 441, 3, 2, 2, 2, 443, 444, 3, 2, 2,
	2, 444, 134, 3, 2, 2, 2, 445, 446, 7, 50, 2, 2, 446, 447, 7, 122, 2, 2,
	447, 449, 3, 2, 2, 2, 448, 450, 9, 7, 2, 2, 449, 448, 3, 2, 2, 2, 450,
	451, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 136,
	3, 2, 2, 2, 453, 454, 7, 50, 2, 2, 454, 458, 9, 8, 2, 2, 455, 457, 9, 9,
	2, 2, 456, 455, 3, 2, 2, 2, 457, 460, 3, 2, 2, 2, 458, 456, 3, 2, 2, 2,
	458, 459, 3, 2, 2, 2, 459, 138, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 461,
	465, 7, 36, 2, 2, 462, 464, 5, 141, 71, 2, 463, 462, 3, 2, 2, 2, 464, 467,
	3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 468, 3, 2,
	2, 2, 467, 465, 3, 2, 2, 2, 468, 469, 7, 36, 2, 2, 469, 140, 3, 2, 2, 2,
	470, 473, 5, 143, 72, 2, 471, 473, 10, 10, 2, 2, 472, 470, 3, 2, 2, 2,
	472, 471, 3, 2, 2, 2, 473, 142, 3, 2, 2, 2, 474, 475, 7, 94, 2, 2, 475,
	488, 9, 11, 2, 2, 476, 477, 7, 94, 2, 2, 477, 478, 7, 119, 2, 2, 478, 479,
	3, 2, 2, 2, 479, 481, 7, 125, 2, 2, 480, 482, 5, 145, 73, 2, 481, 480,
	3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 483, 484, 3, 2,
	2, 2, 484, 485, 3, 2, 2, 2, 485, 486, 7, 127, 2, 2, 486, 488, 3, 2, 2,
	2, 487, 474, 3, 2, 2, 2, 487, 476, 3, 2, 2, 2, 488, 144, 3, 2, 2, 2, 489,
	490, 9, 12, 2, 2, 490, 146, 3, 2, 2, 2, 491, 493, 9, 13, 2, 2, 492, 491,
	3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 492, 3, 2, 2, 2, 494, 495, 3, 2,
	2, 2, 495, 496, 3, 2, 2, 2, 496, 497, 8, 74, 2, 2, 497, 148, 3, 2, 2, 2,
	498, 500, 9, 14, 2, 2, 499, 498, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501,
	499, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 504,
	8, 75, 2, 2, 504, 150, 3, 2, 2, 2, 505, 506, 7, 49, 2, 2, 506, 507, 7,
	44, 2, 2, 507, 512, 3, 2, 2, 2, 508, 511, 5, 151, 76, 2, 509, 511, 11,
	2, 2, 2, 510, 508, 3, 2, 2, 2, 510, 509, 3, 2, 2, 2, 511, 514, 3, 2, 2,
	2, 512, 513, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 513, 515, 3, 2, 2, 2, 514,
	512, 3, 2, 2, 2, 515, 516, 7, 44, 2, 2, 516, 517, 7, 49, 2, 2, 517, 518,
	3, 2, 2, 2, 518, 519, 8, 76, 2, 2, 519, 152, 3, 2, 2, 2, 520, 521, 7, 49,
	2, 2, 521, 522, 7, 49, 2, 2, 522, 526, 3, 2, 2, 2, 523, 525, 10, 14, 2,
	2, 524, 523, 3, 2, 2, 2, 525, 528, 3, 2, 2, 2, 526, 524, 3, 2, 2, 2, 526,
	527, 3, 2, 2, 2, 527, 529, 3, 2, 2, 2, 528, 526, 3, 2, 2, 2, 529, 530,
	8, 77, 2, 2, 530, 154, 3, 2, 2, 2, 20, 2, 412, 416, 420, 426, 435, 443,
	451, 458, 465, 472, 483, 487, 494, 501, 510, 512, 526, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "','", "'{'", "'}'", "':'", "'['", "']'", "'<->'", "'='", "'||'",
	"'&&'", "'.'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'&'", "'!'", "'<-'", "'?'", "", "'as'", "'as?'",
	"'('", "')'", "'transaction'", "'struct'", "'resource'", "'contract'",
	"'interface'", "'fun'", "'event'", "'emit'", "'pre'", "'post'", "'priv'",
	"'pub'", "'pub(set)'", "'return'", "'break'", "'continue'", "'let'", "'var'",
	"'if'", "'else'", "'while'", "'true'", "'false'", "'nil'", "'import'",
	"'from'", "'create'", "'destroy'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Ampersand", "Negate", "Move", "Optional", "NilCoalescing",
	"Downcasting", "FailableDowncasting", "OpenParen", "CloseParen", "Transaction",
	"Struct", "Resource", "Contract", "Interface", "Fun", "Event", "Emit",
	"Pre", "Post", "Priv", "Pub", "PubSet", "Return", "Break", "Continue",
	"Let", "Var", "If", "Else", "While", "True", "False", "Nil", "Import",
	"From", "Create", "Destroy", "Identifier", "DecimalLiteral", "BinaryLiteral",
	"OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral",
	"WS", "Terminator", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "Equal", "Unequal", "Less", "Greater", "LessEqual",
	"GreaterEqual", "Plus", "Minus", "Mul", "Div", "Mod", "Ampersand", "Negate",
	"Move", "Optional", "NilCoalescing", "Downcasting", "FailableDowncasting",
	"OpenParen", "CloseParen", "Transaction", "Struct", "Resource", "Contract",
	"Interface", "Fun", "Event", "Emit", "Pre", "Post", "Priv", "Pub", "PubSet",
	"Return", "Break", "Continue", "Let", "Var", "If", "Else", "While", "True",
	"False", "Nil", "Import", "From", "Create", "Destroy", "Identifier", "IdentifierHead",
	"IdentifierCharacter", "DecimalLiteral", "BinaryLiteral", "OctalLiteral",
	"HexadecimalLiteral", "InvalidNumberLiteral", "StringLiteral", "QuotedText",
	"EscapedCharacter", "HexadecimalDigit", "WS", "Terminator", "BlockComment",
	"LineComment",
}

type CadenceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCadenceLexer(input antlr.CharStream) *CadenceLexer {

	l := new(CadenceLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Cadence.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CadenceLexer tokens.
const (
	CadenceLexerT__0                 = 1
	CadenceLexerT__1                 = 2
	CadenceLexerT__2                 = 3
	CadenceLexerT__3                 = 4
	CadenceLexerT__4                 = 5
	CadenceLexerT__5                 = 6
	CadenceLexerT__6                 = 7
	CadenceLexerT__7                 = 8
	CadenceLexerT__8                 = 9
	CadenceLexerT__9                 = 10
	CadenceLexerT__10                = 11
	CadenceLexerT__11                = 12
	CadenceLexerEqual                = 13
	CadenceLexerUnequal              = 14
	CadenceLexerLess                 = 15
	CadenceLexerGreater              = 16
	CadenceLexerLessEqual            = 17
	CadenceLexerGreaterEqual         = 18
	CadenceLexerPlus                 = 19
	CadenceLexerMinus                = 20
	CadenceLexerMul                  = 21
	CadenceLexerDiv                  = 22
	CadenceLexerMod                  = 23
	CadenceLexerAmpersand            = 24
	CadenceLexerNegate               = 25
	CadenceLexerMove                 = 26
	CadenceLexerOptional             = 27
	CadenceLexerNilCoalescing        = 28
	CadenceLexerDowncasting          = 29
	CadenceLexerFailableDowncasting  = 30
	CadenceLexerOpenParen            = 31
	CadenceLexerCloseParen           = 32
	CadenceLexerTransaction          = 33
	CadenceLexerStruct               = 34
	CadenceLexerResource             = 35
	CadenceLexerContract             = 36
	CadenceLexerInterface            = 37
	CadenceLexerFun                  = 38
	CadenceLexerEvent                = 39
	CadenceLexerEmit                 = 40
	CadenceLexerPre                  = 41
	CadenceLexerPost                 = 42
	CadenceLexerPriv                 = 43
	CadenceLexerPub                  = 44
	CadenceLexerPubSet               = 45
	CadenceLexerReturn               = 46
	CadenceLexerBreak                = 47
	CadenceLexerContinue             = 48
	CadenceLexerLet                  = 49
	CadenceLexerVar                  = 50
	CadenceLexerIf                   = 51
	CadenceLexerElse                 = 52
	CadenceLexerWhile                = 53
	CadenceLexerTrue                 = 54
	CadenceLexerFalse                = 55
	CadenceLexerNil                  = 56
	CadenceLexerImport               = 57
	CadenceLexerFrom                 = 58
	CadenceLexerCreate               = 59
	CadenceLexerDestroy              = 60
	CadenceLexerIdentifier           = 61
	CadenceLexerDecimalLiteral       = 62
	CadenceLexerBinaryLiteral        = 63
	CadenceLexerOctalLiteral         = 64
	CadenceLexerHexadecimalLiteral   = 65
	CadenceLexerInvalidNumberLiteral = 66
	CadenceLexerStringLiteral        = 67
	CadenceLexerWS                   = 68
	CadenceLexerTerminator           = 69
	CadenceLexerBlockComment         = 70
	CadenceLexerLineComment          = 71
)
