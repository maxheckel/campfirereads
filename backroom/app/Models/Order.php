<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Order extends Model
{
    use HasFactory;

    public function shippingAddress(){
        return $this->belongsTo(Address::class, 'shipping_address_id');
    }

    public function lineItems(){
        return $this->hasMany(LineItem::class);
    }
}
