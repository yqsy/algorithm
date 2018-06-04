jsonParse = function () {

    let json = "";

    let p = 0;

    let escapee = {
        "\"": "\"",
        "\\": "\\",
        "/": "/",
        b: "\b",
        f: "\f",
        n: "\n",
        r: "\r",
        t: "\t"
    };


    let isDigit1To9 = function (c) {
        return c >= "1" && c <= "9";
    };

    let isDigit = function (c) {
        return c >= "0" && c <= "9";
    };

    let error = function (m) {
        throw {
            name: "SyntaxError",
            message: m
        };
    };

    let removeWhite = function () {
        let i = p;
        while (i < json.length && json[i] <= " ") {
            i++;
        }
        p = i;
    };

    let removeACharacter = function (c) {
        if (json.length < 1 || json[p] !== c) {
            error("SyntaxError");
        }
        p++;
    };


    let peekACharacter = function () {
        if (json.length < 1) {
            error("SyntaxError");
        }
        return json[p];
    };


    let getString = function () {
        removeACharacter("\"");

        let string_ = "";

        for (let i = p; i < json.length; i++) {
            if (json[i] === "\"") {
                p = i + 1;
                return string_;
            }

            if (json[i] === "\\") {
                i++;

                if (i >= json.length) {
                    break;
                }

                if (typeof escapee[json[i]] === "string") {
                    string_ += escapee[json[i]];
                } else if (json[i] === "u") {
                    if (i + 4 >= json.length) {
                        break;
                    }

                    let r = parseInt(json.slice(i + 1, i + 5), 16);
                    string_ += String.fromCharCode(r);
                    i += 4;
                } else {
                    break;
                }
            } else {
                string_ += json[i];
            }
        }

        error("SyntaxError");
    };

    let getWord = function (s) {
        if (json.length < s.length) {
            error("SyntaxError");
        }

        let rtn = json.slice(p, p + s.length);

        if (rtn !== s) {
            error("SyntaxError");
        }

        p += s.length;
        return rtn;
    };

    let parseString = function () {
        return getString();
    };

    let parseWord = function (specifyStr, kind) {
        getWord(specifyStr);
        return kind;
    };

    let parseNumber = function () {
        let j = p;

        let jValid = function () {
            return j < json.length;
        };

        if (jValid() && json[j] === "-") {
            j++;
        }

        if (jValid() && json[j] === "0") {
            j++;
        } else {
            if (!jValid() || !isDigit1To9(json[j])) {
                error("SyntaxError");
            }
            j++;
            while (jValid() && isDigit(json[j])) {
                j++
            }
        }

        if (jValid() && json[j] === ".") {
            j++;
            if (!jValid() || !isDigit(json[j])) {
                error("SyntaxError");
            }
            while (jValid() && isDigit(json[j])) {
                j++;
            }
        }

        if (jValid() && (json[j] === 'e' || json[j] === 'E')) {
            j++;
            if (jValid() && (json[j] === '+' || json[j] === '-')) {
                j++
            }
            if (!jValid() || !isDigit(json[j])) {
                error("SyntaxError");
            }
            while (jValid() && isDigit(json[j])) {
                j++;
            }
        }

        let str = json.slice(p, j);
        p = j;
        return parseFloat(str);
    };


    let parseObject = function () {
        removeACharacter("{");

        let obj = {};

        while (1) {
            removeWhite();
            let key = getString();

            removeWhite();
            removeACharacter(":");

            obj[key] = parseValue();

            removeWhite();
            if (peekACharacter() !== ',') {
                break;
            } else {
                removeACharacter(",");
            }
        }

        removeWhite();
        removeACharacter("}");

        return obj;
    };


    let parseArray = function () {
        removeACharacter("[");

        let array = [];

        while (1) {
            array.push(parseValue());

            removeWhite();
            if (peekACharacter() !== ',') {
                break;
            } else {
                removeACharacter(",");
            }
        }

        removeWhite();
        removeACharacter("]");
        return array;
    };


    let parseValue = function () {
        removeWhite();

        let c = peekACharacter();

        switch (c) {
            case "\"":
                return parseString();
            case "n":
                return parseWord("null", null);
            case "t":
                return parseWord("true", true);
            case "f":
                return parseWord("false", false);
            case "[":
                return parseArray();
            case "{":
                return parseObject();
            default:
                return parseNumber();
        }
    };


    return function (text) {
        json = text;
        p = 0;

        return parseValue();
    };

}();


module.exports = {jsonParse: jsonParse};

