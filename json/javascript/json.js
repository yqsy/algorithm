if (typeof MyJson !== "object") {
    MyJson = {};
}

(function () {
    "use script";

    let ch; // current character
    let at; // current idx
    let text;

    let error = function (m) {

        throw {
            name: "SyntaxError",
            message: m,
            at: at,
            text: text
        };
    };

    let value = function () {


    };


    if (typeof MyJson.parse !== "function") {
        MyJson.parse = function (source, reviver) {
            text = source;

            if (typeof reviver === "function") {
                // 调用回调函数
                // 暂时不实现
            } else {

            }
        };
    }

}());

