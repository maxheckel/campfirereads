<?php
namespace App\Services\Payments;

use App\Models\Address;
use App\Models\LineItem;
use App\Models\Order;
use Carbon\Carbon;
use Illuminate\Support\Facades\DB;
use Stripe\Stripe;
use Stripe\StripeClient;

class StripeService implements PaymentService{

    public function __construct(private StripeClient $client){

    }

    public function GetOrderForID(string $id): Order
    {
        $order = Order::where('order_id', $id)->with('shippingAddress', 'lineItems')->first();
        if (!$order){
            $cs = $this->client->checkout->sessions->retrieve($id, [
                'expand'=>['payment_intent']
            ]);
            $lineItems = $this->client->checkout->sessions->allLineItems($id,[
                'expand'=>['data.price.product']
            ]);


            DB::beginTransaction();
            $order = new Order();
            $order->internal_id = $cs->client_reference_id;
            $order->payment_id = $cs->payment_intent->id;
            $order->order_id = $id;
            $order->ordered_on = Carbon::parse($cs->created);
            $shipping = new Address();
            /** @var  $piShipping */

            $shipping->name = $cs->payment_intent->shipping->name;
            $shipping->street1 = $cs->payment_intent->shipping->address->line1;
            $shipping->street2 = $cs->payment_intent->shipping->address->line2;
            $shipping->city = $cs->payment_intent->shipping->address->city;
            $shipping->state = $cs->payment_intent->shipping->address->state;
            $shipping->zip = $cs->payment_intent->shipping->address->postal_code;
            $shipping->country = $cs->payment_intent->shipping->address->country;
            $shipping->save();
            $order->shipping_address_id = $shipping->id;
            $order->save();

            foreach ($lineItems->data as $lineItem){
                $li = new LineItem();
                $li->order_id = $order->id;
                $li->product_id = $lineItem->price->product->id;
                $li->price_id = $lineItem->price->id;
                $li->type = $lineItem->price->metadata->listing_type ?? 'unknown';
                $li->amazon_url = "https://amazon.com/".trim($lineItem->price->product->metadata->amazon_url, '/');
                $li->save();
            }
            DB::commit();
        }
        return $order;
    }

    public function GetNewOrders(): array
    {
        return [];
    }
}
