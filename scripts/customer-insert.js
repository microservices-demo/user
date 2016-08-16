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
    "password": "b5040cba4bb01c0a9570d261fd031d6d67e384e2",
    "addresses": [ObjectId("57a98d98e4b00679b4a830ad")],
    "cards": [ObjectId("57a98d98e4b00679b4a830ae")]
});
insert_customer({
    "_id": ObjectId("57a98d98e4b00679b4a830b2"),
    "firstName": "User",
    "lastName": "Name",
    "username": "user",
    "password": "d2769601c686ed2c4a16cccdb3c7104e1fe2a5ef",
    "addresses": [ObjectId("57a98d98e4b00679b4a830b0")],
    "cards": [ObjectId("57a98d98e4b00679b4a830b1")]
});
insert_customer({
    "_id": ObjectId("57a98d98e4b00679b4a830b5"),
    "firstName": "User1",
    "lastName": "Name1",
    "username": "user1",
    "password": "5cd681c7e60178b4d566bdff96162475a53dd41a",
    "addresses": [ObjectId("57a98d98e4b00679b4a830b3")],
    "cards": [ObjectId("57a98d98e4b00679b4a830b4")]
});
print("_______CUSTOMER DATA_______");
db.customers.find().forEach(get_results);
print("______END CUSTOMER DATA_____");
