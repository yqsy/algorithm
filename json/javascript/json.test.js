let expect = require('chai').expect;
let _ = require('lodash');


let jsonParse = require('./json').jsonParse;

let isDoubleEqual = function (f1, f2) {
    return Math.abs(f1 - f2) < 0.000001;
};


describe("test number", function () {
    it('', function () {
        expect(isDoubleEqual(1.0, jsonParse('1 '))).to.be.equal(true);
        expect(isDoubleEqual(0.0, jsonParse("-0 "))).to.be.equal(true);
        expect(isDoubleEqual(0.0, jsonParse("-0.0 "))).to.be.equal(true);
        expect(isDoubleEqual(1, jsonParse("1 "))).to.be.equal(true);

        expect(isDoubleEqual(-1.0, jsonParse("-1 "))).to.be.equal(true);
        expect(isDoubleEqual(1.5, jsonParse("1.5 "))).to.be.equal(true);
        expect(isDoubleEqual(-1.5, jsonParse("-1.5 "))).to.be.equal(true);
        expect(isDoubleEqual(3.1416, jsonParse("3.1416 "))).to.be.equal(true);
        expect(isDoubleEqual(1E10, jsonParse("1E10 "))).to.be.equal(true);
        expect(isDoubleEqual(1e10, jsonParse("1e10 "))).to.be.equal(true);
        expect(isDoubleEqual(1E+10, jsonParse("1E+10 "))).to.be.equal(true);
        expect(isDoubleEqual(1E-10, jsonParse("1E-10 "))).to.be.equal(true);
        expect(isDoubleEqual(-1E10, jsonParse("-1E10 "))).to.be.equal(true);
        expect(isDoubleEqual(-1e10, jsonParse("-1e10 "))).to.be.equal(true);
        expect(isDoubleEqual(-1E+10, jsonParse("-1E+10 "))).to.be.equal(true);
        expect(isDoubleEqual(-1E-10, jsonParse("-1E-10 "))).to.be.equal(true);
        expect(isDoubleEqual(1.234E+10, jsonParse("1.234E+10 "))).to.be.equal(true);
        expect(isDoubleEqual(1.234E-10, jsonParse("1.234E-10 "))).to.be.equal(true);
        expect(isDoubleEqual(0.0, jsonParse("1e-10000 "))).to.be.equal(true);
        /* must underflow */

        expect(isDoubleEqual(1.0000000000000002, jsonParse("1.0000000000000002 "))).to.be.equal(true);
        /* the smallest number > 1 */
        expect(isDoubleEqual(4.9406564584124654e-324, jsonParse("4.9406564584124654e-324 "))).to.be.equal(true);
        /* minimum denormal */
        expect(isDoubleEqual(-4.9406564584124654e-324, jsonParse("-4.9406564584124654e-324 "))).to.be.equal(true);
        expect(isDoubleEqual(2.2250738585072009e-308, jsonParse("2.2250738585072009e-308 "))).to.be.equal(true);
        /* Max subnormal double */
        expect(isDoubleEqual(-2.2250738585072009e-308, jsonParse("-2.2250738585072009e-308 "))).to.be.equal(true);
        expect(isDoubleEqual(2.2250738585072014e-308, jsonParse("2.2250738585072014e-308 "))).to.be.equal(true);
        /* Min normal positive double */
        expect(isDoubleEqual(-2.2250738585072014e-308, jsonParse("-2.2250738585072014e-308 "))).to.be.equal(true);
        expect(isDoubleEqual(1.7976931348623157e+308, jsonParse("1.7976931348623157e+308 "))).to.be.equal(true);
        /* Max double */
        expect(isDoubleEqual(-1.7976931348623157e+308, jsonParse("-1.7976931348623157e+308 "))).to.be.equal(true);
    });
});

describe("test string", function () {
    it('', function () {
        expect(jsonParse('""')).to.be.equal("");
        expect(jsonParse('"Hello"')).to.be.equal("Hello");
        expect(jsonParse('"Hello\\nWorld"')).to.be.equal("Hello\nWorld");
        expect(jsonParse('"\\" \\\\ \\/ \\b \\f \\n \\r \\t"')).to.be.equal("\" \\ / \b \f \n \r \t");
        expect(jsonParse('"Hello\\u0000World"')).to.be.equal("Hello\u0000World");
        expect(jsonParse('"\\u0024"')).to.be.equal("\x24");
        //expect(jsonParse('"\\u20AC"')).to.be.equal("\xE2\x82\xAC");
        expect(jsonParse('"你好吗"')).to.be.equal("你好吗");
    });
});

describe("test object", function () {
    it('', function () {
        let json = `
 {
 "n" : null  ,
"f" : false ,
"t" : true,
"i" : 123,
"s" : "abc", 
"a" : [ 1, 2, 3],
"o" : { "1" : 1, "2" : 2, "3" : 3}
}`;

        let obj = jsonParse(json);
        expect(obj["n"] === null).to.be.equal(true);
        expect(obj["f"] === false).to.be.equal(true);
        expect(obj["t"] === true).to.be.equal(true);
        expect(isDoubleEqual(obj["i"], 123)).to.be.equal(true);
        expect(obj["s"] === "abc").to.be.equal(true);
        expect(isDoubleEqual(obj["a"][2], 3)).to.be.equal(true);
        expect(isDoubleEqual(obj["o"]["2"], 2)).to.be.equal(true);
    });
});

describe("test simple kind", function () {
    it('', function () {
        expect(jsonParse("null") === null).to.be.equal(true);
        expect(jsonParse("true") === true).to.be.equal(true);
        expect(jsonParse(`"shit"`) === "shit").to.be.equal(true);
        expect(isDoubleEqual(jsonParse(`123456`),123456)).to.be.equal(true);
        expect(isDoubleEqual(jsonParse(`[1,2,3]`)[0],1)).to.be.equal(true);
    });
});
