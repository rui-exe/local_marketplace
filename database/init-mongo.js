db = db.getSiblingDB('marketplace');

db.createCollection('user');
db.user.createIndex({ "email": 1 }, { unique: true }); // Ensure email is unique
db.user.createIndex({ "phone": 1 }, { unique: true }); // Ensure phone is unique


