<template>


    <div class="text-center p-8 bg-center bg-cover	" :style="{'background-image': 'url('+data.bgImage+')'}">
      <div class="z-10 w-full  relative left-0 p-8">
        <div class="trapezoid text-np-yellow p-10 text-center w-full  relative">

          <Header :invert="true" class="text-white z-10 block mx-auto justify-center relative" :width-override="'w-10'"
                  :text="'Search results for \''+query+'\''" :icon-path="'/media/arrowhead.svg'"></Header>
          <span class="text-np-yellow block relative z-10 md:text-2xl tracking-widest md:mt-4"
                style="font-family: Freehand">
            <template v-if="!data.loading">
              <template v-if="data.page > 1">
                Page {{data.page}} -
              </template>
              {{ data.metadata.totalItems }} Found
            </template>
            <template v-else>
              Searching...
            </template>

        </span>
        </div>

      </div>
    </div>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">

    <div v-if="data.loading">
      <div v-for="i in new Array(10)" class="grid grid-cols-[20%_80%] gap-4 my-4 border-b-1 border-b-gray-200 border-b py-4 border-box">
        <ShimmerBox class="w-full h-[220px]"/>
        <div>
          <ShimmerBox class="h-8 w-40 rounded-full"></ShimmerBox>
          <ShimmerBox class="w-80 mt-4 h-4 rounded-full"></ShimmerBox>
          <ShimmerBox class="w-80 mt-4 h-4 rounded-full"></ShimmerBox>
          <ShimmerBox class="w-60 mt-4 h-4 rounded-full"></ShimmerBox>
        </div>
      </div>
    </div>
    <div v-else class="grid grid-cols-2 sm:grid-cols-3 md:block gap-8">
      <BookSearchResult :book="book" v-for="book in data.books"></BookSearchResult>
    </div>

    <div class="mt-8 block">
      <Button @click="previousPage()" class="float-left" v-if="data.page > 1" :text="'Previous Page'"></Button>
      <Button @click="nextPage()" class="float-right" v-if="showNextPage()" :text="'Next Page'"></Button>
    </div>



  </div>

</template>

<script setup>

import { useRoute } from 'vue-router'
import {computed, reactive, watch} from "vue";
import Header from "../components/Header.vue";
import BookSearchResult from "../components/BookSearchResult.vue";
import ShimmerBox from "../components/ShimmerBox.vue";
import {decimalHash, bookHref} from "../services/utils.js";
import Button from "../components/Button.vue";
import router from "../router";

const route = useRoute()
const query = route.query.query
const page = parseInt(route.query.page)
const data = reactive({
  books: [],
  loading: true,
  bgImage: "/media/covers/" + (decimalHash(query) + "").split('.')[1].slice(0, 2) + ".jpg",
  metadata: {},
  page: page
})

fetch(import.meta.env.VITE_API_HOST + "search?page="+page+"&query="+query)
    .then((response) => response.json())
    .then((resp) => {
      data.loading = false
      data.metadata = resp
      data.books = resp.items.filter((b) => bookHref(b) !== undefined)
    });

function showNextPage(){
  if((data.page+1)*20 > data.metadata.totalItems){
    return false
  }
  return true
}

watch(() => data.page, () => {
  data.loading = true;

  fetch(import.meta.env.VITE_API_HOST + "search?page="+data.page+"&query="+query)
      .then((response) => response.json())
      .then((resp) => {
        data.loading = false
        data.metadata = resp
        data.books = resp.items.filter((b) => bookHref(b) !== undefined)
      });

});


watch(
    () => router.currentRoute,
    () => {
      console.log('new search')
    }
)

function previousPage(){
  if(!data.page){
    data.page = 1
  }
  data.page = parseInt(data.page) - 1
  router.push({name:'search', query:{query: query, page: data.page}})
  document.body.scrollTop = document.documentElement.scrollTop = 0;
}

function nextPage(){
  if(!data.page){
    data.page = 1
  }
  data.page = parseInt(data.page) + 1
  router.push({name:'search', query:{query: query, page: data.page}})
  document.body.scrollTop = document.documentElement.scrollTop = 0;
}
</script>

<style scoped>

.trapezoid {
  display: inline-block;

  position: relative;
}

.trapezoid::before,
.trapezoid::after {
  content: "";
  position: absolute;
  top: 0;
  border: 4px solid #F0E9CF;
  bottom: 0;
  width: 60%;
  background: #221B1C;
  transform-origin: top;
  box-sizing: border-box;
  z-index: 0;
}

.trapezoid::before {
  border-radius: 30px 0 0 30px;
  border-right: 0;
  left: 0;
  transform: skew(8deg);
}

.trapezoid::after {
  border-radius: 0 30px 30px 0;
  border-left: 0;
  right: 0;
  transform: skew(-8deg);
}

</style>