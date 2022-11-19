<template>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">
    <h1 class="text-5xl font-bold uppercase">Page not found!</h1>
    The resource you were looking for is not found.
    <span class="material-symbols-outlined text-8xl font-light text-np-yellow opacity-30 block">
      local_fire_department
      </span>
    <BookList class="mt-20" title="Try some of our popular titles instead" :endpoint="'popular'"  :category-link="'/browse/popular'"></BookList>
  </div>
</template>

<script setup>

import BookList from "../components/BookList.vue";
</script>

<style scoped>

</style>