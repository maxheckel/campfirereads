<?php
namespace App\Services\Payments;


use App\Models\Order;

interface PaymentService{
    public function GetOrderForID(string $id):Order;
    public function GetNewOrders():array;
}
