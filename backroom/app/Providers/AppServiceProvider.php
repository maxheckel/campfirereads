<?php

namespace App\Providers;

use App\Services\Payments\PaymentService;
use App\Services\Payments\StripeService;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        $this->app->bind(PaymentService::class, StripeService::class);
        $this->app->bind(StripeService::class, function ($a){
            return new StripeService(new \Stripe\StripeClient(env('STRIPE_SECRET_KEY')));
        });
    }
}
