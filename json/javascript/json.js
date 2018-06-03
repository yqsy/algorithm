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

    let error = function (m) {
        throw {
            name: "SyntaxError",
            message: m
        };
    };

    let removeWhite = function () {
        let i = 0;
        while (i < json.length && json[i] <= " ") {
            i++;
        }
        p = i;
    };

    let removeACharacter = function (c) {
        if (json.length < 1 || json[0] !== c) {
            error("SyntaxError");
        }
        p++;
    };


    let peekACharacter = function () {
        if (json.length < 1) {
            error("SyntaxError");
        }
        return json[0];
    };


    let getString = function () {
        removeACharacter("\"");

        let string = "";

        for (let i = 0; i < json.length; i++) {
            if (json[i] === "\"") {
                p = i + 1;
                return string;
            }

            if (json[i] === "\\") {
                i++;

                if (i >= json.length) {
                    break;
                }

                if (typeof escapee[json[i]] === "string") {
                    string += escapee[json[i]];
                } else if (json[i] === "u") {
                    if (i + 4 >= json.length) {
                        break;
                    }

                    let r = parseInt(json.slice(i + 1, i + 5), 16);
                    string += String.fromCharCode(r);
                    i += 4;
                } else {
                    break;
                }

            }
        }

        error("SyntaxError");
    };

    return function (text) {
        json = text;
        p = 0;


    };

}();





