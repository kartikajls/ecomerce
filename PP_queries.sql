

-- 1. USER REPORTS: Daftar akun beserta identitas profilnya
SELECT u.id AS user_id, u.email, u.role, COALESCE(p.full_name, '-') AS full_name, COALESCE(p.phone, '-') AS phone
FROM users u
LEFT JOIN profiles p ON u.id = p.user_id
ORDER BY u.id ASC;

-- 2. STOCK REPORTS: Status ketersediaan stok produk pakaian
SELECT sku, name, size, color, stock,
       CASE 
           WHEN stock = 0 THEN 'HABIS' 
           WHEN stock < 5 THEN 'MENIPIS' 
           ELSE 'AMAN' 
       END AS status_stok
FROM products
ORDER BY stock ASC;

-- 3. ORDER REPORTS: Riwayat transaksi pesanan yang masuk
SELECT o.id AS order_id, u.email AS pembeli, o.total_amount, o.status, DATE_FORMAT(o.order_date, '%Y-%m-%d %H:%i') AS waktu_transaksi
FROM orders o
JOIN users u ON o.user_id = u.id
ORDER BY o.order_date DESC;