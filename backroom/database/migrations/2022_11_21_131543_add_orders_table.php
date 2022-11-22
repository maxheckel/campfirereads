<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('orders', function (Blueprint $table) {
            $table->id();
            $table->timestamps();
            $table->string('order_id');
            $table->string('internal_id');
            $table->string('payment_id');
            $table->string('order_provider')->default('stripe');
            $table->timestamp('ordered_on');
            $table->boolean('fulfilled')->default(false);
            $table->foreignIdFor(\App\Models\Address::class, 'shipping_address_id');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('orders');
    }
};
