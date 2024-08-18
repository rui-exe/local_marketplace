// Connect to the marketplace database
db = db.getSiblingDB('marketplace');

// Create the Users collection and set indexes
db.createCollection('users');
db.users.createIndex({ "email": 1 }, { unique: true }); // Ensure email is unique
db.users.insertMany([
  {
    username: "john_doe",
    email: "john@example.com",
    password: "hashed_password_here", // Note: Passwords should be hashed in your application code
    role: "seller",
    createdAt: new Date(),
  },
  {
    username: "jane_doe",
    email: "jane@example.com",
    password: "hashed_password_here",
    role: "buyer",
    createdAt: new Date(),
  }
]);

// Create the Products collection and set indexes
db.createCollection('products');
db.products.createIndex({ "name": "text", "description": "text" }); // Full-text search index
db.products.createIndex({ "category": 1 });
db.products.createIndex({ "sellerId": 1 });
db.products.insertMany([
  {
    name: "Handcrafted Wooden Chair",
    description: "A beautiful handcrafted wooden chair made from oak.",
    price: 150.00,
    category: "Furniture",
    images: ["chair1.jpg", "chair2.jpg"],
    sellerId: db.users.findOne({ username: "john_doe" })._id,
    location: "New York",
    createdAt: new Date(),
  },
  {
    name: "Organic Honey",
    description: "Pure organic honey from local farms.",
    price: 10.00,
    category: "Food",
    images: ["honey.jpg"],
    sellerId: db.users.findOne({ username: "john_doe" })._id,
    location: "California",
    createdAt: new Date(),
  }
]);

// Create the Orders collection
db.createCollection('orders');
db.orders.createIndex({ "buyerId": 1 });
db.orders.createIndex({ "sellerId": 1 });
db.orders.insertOne({
  buyerId: db.users.findOne({ username: "jane_doe" })._id,
  sellerId: db.users.findOne({ username: "john_doe" })._id,
  productId: db.products.findOne({ name: "Handcrafted Wooden Chair" })._id,
  quantity: 1,
  totalAmount: 150.00,
  status: "Pending",
  createdAt: new Date(),
});

// Optional: Create other collections like Messages, Reviews, etc.
