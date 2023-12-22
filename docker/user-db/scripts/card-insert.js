function get_results(result) {
    print(tojson(result));
}

function insert_card(object) {
    print(db.cards.insert(object));
}

insert_card({
    "_id": ObjectId("57a98d98e4b00679b4a830ae"),
    "longNum": "4833571284527541",
    "expires": "08/19",
    "ccv": "678"
});
insert_card({
    "_id": ObjectId("57a98d98e4b00679b4a830b1"),
    "longNum": "5186025451630844",
    "expires": "08/19",
    "ccv": "958"
});
insert_card({
    "_id": ObjectId("57a98d98e4b00679b4a830b4"),
    "longNum": "375519564445134",
    "expires": "08/19",
    "ccv": "280"
});
insert_card({
    "_id": ObjectId("57a98ddce4b00679b4a830d2"),
    "longNum": "4511605861436",
    "expires": "04/16",
    "ccv": "432"
});

print("________CARD DATA_______");
db.cards.find().forEach(get_results);
print("______END CARD DATA_____");


