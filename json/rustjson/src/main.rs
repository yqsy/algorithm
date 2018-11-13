#[cfg(not(feature = "preserve_order"))]
use std::collections::{btree_map, BTreeMap};

#[derive(Clone, PartialEq)]
pub struct Number {
    n: N,
}

#[derive(PartialOrd, Eq, Ord)]
#[stable(feature = "rust1", since = "1.0.0")]
pub struct String {
    vec: Vec<u8>,
}

pub struct Map<K, V> {
    map: MapImpl<K, V>
}

#[cfg(not(feature = "preserve_order"))]
type MapImpl<K, V> = BTreeMap<K, V>;


#[derive(Clone, PartialEq)]
pub enum Value {
    Null,
    Bool(bool),
    Number(Number),
    String(String),
    Array(Vec<Value>),
    Object(Map<String, Value>),
}




fn main() {}
