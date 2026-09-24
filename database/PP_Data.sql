
-- Akun Admin & Customer (Password dummy: admin123 & user123)
INSERT INTO users (email, password, role) VALUES 
('admin@store.com', 'admin123', 'admin');

INSERT INTO profiles (user_id, full_name, phone, address) VALUES 
(1, 'Admin Gudang Fesyen', '081111111111', 'Jakarta Selatan, DKI Jakarta'),
(2, 'Imam Pembeli', '082222222222', 'Bandung, Jawa Barat');

DELETE FROM users WHERE id IN (1, 2);

INSERT INTO users (email, password, role) VALUES 
('admin@mail.com', 'admin456', 'admin'),
('paul_pogba@mail.com', 'user456', 'customer');

INSERT INTO profiles (user_id, full_name, phone, address) VALUES 
(1, 'admin toko clothing', '081234567891', 'Jakarta Selatan, DKI Jakarta'),
(2, 'Paul bin Pogba', '082111234577', 'Bandung, Jawa Barat');

-- Kategori Pakaian
INSERT INTO categories (name) VALUES 
('Atasan Pria'), ('Outerwear'), ('Celana');

-- 1. Matikan sementara pengecekan relasi antar tabel
SET FOREIGN_KEY_CHECKS = 0;

-- 2. Kosongkan semua tabel & reset angka AUTO_INCREMENT kembali ke 1
TRUNCATE TABLE payments;
TRUNCATE TABLE order_details;
TRUNCATE TABLE orders;
TRUNCATE TABLE products;
TRUNCATE TABLE profiles;
TRUNCATE TABLE users;
TRUNCATE TABLE categories;

-- 3. Nyalakan kembali pengecekan relasi
SET FOREIGN_KEY_CHECKS = 1;

-- 4. Masukkan data baru (Users)
INSERT INTO users (email, password, role) VALUES 
('admin@mail.com', 'admin456', 'admin'),
('paul_pogba@mail.com', 'user456', 'customer');

-- 5. Masukkan data baru (Profiles)
INSERT INTO profiles (user_id, full_name, phone, address) VALUES 
(1, 'admin toko clothing', '081234567891', 'Jakarta Selatan, DKI Jakarta'),
(2, 'Paul bin Pogba', '082111234577', 'Bandung, Jawa Barat');

-- 6. Masukkan data baru (Kategori Pakaian)
INSERT INTO categories (name) VALUES 
('Atasan Pria'), ('Outerwear'), ('Celana');

-- Katalog Produk
INSERT INTO products (category_id, sku, name, description, size, color, price, stock) VALUES 
(1, 'TSH-BLK-L', 'Kaos Polos Oversize', 'Kaos katun bambu super adem 24s', 'L', 'Hitam', 120000.00, 20),
(1, 'TSH-WHT-M', 'Kaos Polos Heavyweight', 'Kaos bahan tebal tidak terawang 16s', 'M', 'Putih', 135000.00, 15),
(2, 'JKT-DNM-XL', 'Jaket Denim Vintage', 'Jaket jeans bertekstur klasik klasik', 'XL', 'Biru Navy', 350000.00, 5),
(3, 'CHN-KRM-L', 'Celana Chino Slimfit', 'Celana formal kasual bahan stretch', 'L', 'Krem', 210000.00, 10);

select * from payments;
select * from categories;
select * from products;
select * from profiles;
select * from users u ;
select * from orders o ;



