<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-10">
    <Header class="justify-center " text="Why does this exist?" icon-path="/media/pine.svg"></Header>
    <span class="text-2xl block">Our Story</span>
    <div class="my-4 mb-8">
    Campfire Reads was created in 2022 after our founder, Shane, went camping with his fiance.  After arriving home from a lovely trip Shane picked up the book he had been reading around the fire and discovered something.  It smelled phenomenal.  Immediately struck with inspiration he called Max and said those two two fateful words "Smoked books" to which Max responded with two more, arguably equally as fateful, words..."What's that?".
    </div>
    <div class="grid grid-cols-2 gap-4">
      <div class="flex items-center justify-center">
        <AboutCard class="w-2/3 lg:w-1/2" name="Shane Smith" image="/media/shane.png" title="Founder & CSO (Chief Smoke Officer)" description="Favorite book? I don't have a favorite book, I love all books equally....Harry potter and the prisoner of Azkaban"></AboutCard>
      </div>
      <div class="flex items-center justify-center">
        <AboutCard class="w-2/3 lg:w-1/2" name="Max Heckel" image="/media/max.png" title="Founder & CCO (Chief Camping Officer)" description="I never read unsmoked books now.  When someone hands me a book that smells like paper I'm just like...gross."></AboutCard>
      </div>
    </div>
  </div>
</template>

<script setup>
import Header from "../components/Header.vue";
import AboutCard from "../components/AboutCard.vue";

</script>

<style scoped>

</style>