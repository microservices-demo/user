function get_results(result) {
    print(tojson(result));
}

function insert_customer(object) {
    print(db.customers.insert(object));
}

insert_customer({
    "_id": ObjectId("57a98d98e4b00679b4a830af"),
    "firstName": "Eve",
    "lastName": "Berger",
    "username": "Eve_Berger",
    "password": "f0b6dab1610562b77078081c88c334f934dd85d6",
    "salt": "c748112bc027878aa62812ba1ae00e40ad46d497",
    "addresses": [ObjectId("57a98d98e4b00679b4a830ad")],
    "cards": [ObjectId("57a98d98e4b00679b4a830ae")]
});
//pass test1
insert_customer({
    "_id": ObjectId("57a98d98e4b00679b4a830b2"),
    "firstName": "User",
    "lastName": "Name",
    "username": "user",
    "password": "737a21044f994ca25906702c27157ce3f7633f76",
    "salt": "6c1c6176e8b455ef37da13d953df971c249d0d8e",
    "addresses": [ObjectId("57a98d98e4b00679b4a830b0")],
    "cards": [ObjectId("57a98d98e4b00679b4a830b1")]
});
//pass test2
insert_customer({
    "_id": ObjectId("57a98d98e4b00679b4a830b5"),
    "firstName": "User1",
    "lastName": "Name1",
    "username": "user1",
    "password": "9ffa074f5129ee82ebc44b5229e5e9b9915fe519",
    "salt": "bd832b0e10c6882deabc5e8e60a37689e2b708c2",
    "addresses": [ObjectId("57a98d98e4b00679b4a830b3")],
    "cards": [ObjectId("57a98d98e4b00679b4a830b4")]
});
//pass test3
print("_______CUSTOMER DATA_______");
db.customers.find().forEach(get_results);
print("______END CUSTOMER DATA_____");
