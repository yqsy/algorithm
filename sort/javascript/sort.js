function quickSort(array, first, last) {
    if (first < last) {
        let mid = partition(array, first, last);
        quickSort(array, first, mid - 1);
        quickSort(array, mid + 1, last);
    }
}

function partition(array, first, last) {
    let cur = array[first];
    let p = first + 1;
    for (let i = first + 1; i <= last; i++) {
        if (array[i] < cur) {
            swap(array, i, p);
            p++;
        }
    }
    swap(array, first, p - 1);
    return p - 1;
}


function swap(array, i, j) {
    let temp = array[i];
    array[i] = array[j];
    array[j] = temp;
}


module.exports = {quickSort: quickSort};
