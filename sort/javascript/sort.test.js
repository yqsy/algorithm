let expect = require('chai').expect;
let _ = require('lodash');

let quickSort = require('./sort').quickSort;

function arraysEqual(a, b) {
    if (a === b) return true;
    if (a == null || b == null) return false;
    if (a.length !== b.length) return false;

    for (let i = 0; i < a.length; ++i) {
        if (a[i] !== b[i]) return false;
    }
    return true;
}

describe('test quickSort1', function () {
    it('', function () {
        let array = [5, 7, 6, 4, 8, 9, 2, 1, 3];
        quickSort(array, 0, array.length - 1);
        expect(arraysEqual(array, [1, 2, 3, 4, 5, 6, 7, 8, 9])).to.be.equal(true);
    });

    it('', function () {
        let array = [3, 2];
        quickSort(array, 0, array.length - 1);
        expect(arraysEqual(array, [2, 3])).to.be.equal(true);
    });

    it('', function () {
        let array = [3];
        quickSort(array, 0, array.length - 1);
        expect(arraysEqual(array, [3])).to.be.equal(true);
    });
});
