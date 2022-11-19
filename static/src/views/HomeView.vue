<script setup>
import BookList from "./../components/BookList.vue";
import {onMounted, reactive} from "vue";
import ShimmerBox from "../components/ShimmerBox.vue";
import Search from "../components/Search.vue";

const data = reactive({
  bestSellerLists: Object,
  loading: true
})
onMounted(() => {
  fetch(import.meta.env.VITE_API_HOST + "bestsellers")
      .then((response) => response.json())
      .then((resp) => {
        data.bestSellerLists = resp.lists
        data.loading = false
      });
})
</script>

<template class="min-h-screen">
  <div class="mb-0 border-b-np-yellow-200 border-b-2 relative pt-10 bg-center bg-cover"
       style="background: url('/media/covers/01.jpg')">
    <div class="w-full overflow-hidden max-h-[500px] h-[500px] absolute top-0 z-0">
      <video style="position: absolute; z-index: 0; top: 0; left: 0; width: 100%;" autoplay muted loop>
        <source src="/media/campfire-hq.mp4" type="video/mp4">
        Your browser does not support the video tag.
      </video>

    </div>


    <div class="z-10 w-full lg:w-1/2 lg:ml-[10%] relative left-0 p-8">
      <div class="top relative z-0"></div>
      <div class="trapezoid text-np-yellow z-10 mt-4 md:mt-0">
        <img src="/media/arrow.webp" class="absolute w-20 right-20 z-10 -top-10">
        <div class="z-10 text-4xl relative px-10 -mt-2 lg:-mt-6 float-left" style="font-family: Freehand">
          Welcome To
        </div>
        <div class="relative z-10 p-10 uppercase text-6xl text-center tracking-wider">
          Campfire Reads
          <div class="text-3xl normal-case mt-2" style="font-family: Freehand">
            Premium Smoked Books
          </div>
        </div>

        <div
            class="w-full flex center justify-center items-center -mt-5 md:mt-0 p-2 bg-np-yellow text-center relative z-10 text-[#221B1C] text-sm md:text-lg rounded">
          Read like you're around the campfire
          <div class="inline-block rounded-full mx-2 w-2 h-2 bg-black"></div>
          From the comfort of your own home.
        </div>
      </div>
    </div>
  </div>

  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 md:mt-16  px-4">
    <div class="my-20 md:grid grid-cols-2 gap-8">
      <div>
        <div class="mt-8 md:mt-4 text-4xl">Increase the <span class="bg-np-dark-brown text-np-yellow px-2 rounded-lg py-1 inline-block">emotional intensity</span>
          of your reading
        </div>
        <div class="quote">
          <blockquote class="p-4 my-4 bg-np-yellow-200 border-l-4 border-np-yellow dark:border-gray-500 dark:bg-gray-800">
            <p class="text-xl italic font-medium leading-relaxed text-gray-900 dark:text-white">
              "Smell causes us to time travel in our minds and evoke emotion."
              - Nature.com<br>
              <a href="https://www.nature.com/articles/d41586-022-01626-x" class="text-xs">https://www.nature.com/articles/d41586-022-01626-x</a>
            </p>
          </blockquote>
          Each of our books ships to you smelling of campfire, unlocking instant nostalgia. The smell can last from 6 months
          to years depending on use and material.
        </div>
      </div>
      <div class="hidden md:block">
        <img src="/media/woods-read.jpg" class="rounded rounded-2xl">
      </div>

    </div>

    <Search class="my-8"/>
    <BookList title="Popular Titles" :endpoint="'popular'" :category-link="'/browse/popular'"></BookList>
    <BookList v-if="!data.loading" class="my-20" v-for="list in data.bestSellerLists" :title="list.list.display_name"
              :books="list.books" :category-link="'/browse/'+list.list.list_name_encoded"></BookList>
    <div class="grid-cols-2 md:grid-cols-6 gap-10 grid mt-20" v-if="data.loading">
      <div v-for="i in new Array(48)">
        <ShimmerBox class="w-full h-[250px] bg-gray-200"></ShimmerBox>
      </div>
    </div>
  </div>


</template>

<style scoped lang="scss">

.top {
  width: calc(100% - 50px);
  background-color: #221B1C;
  border-top-left-radius: 50px;
  border-bottom-left-radius: 40px;
  height: 100px;
  margin-bottom: -85px;
  transform: skewY(7deg);
  z-index: 10;
  color: #D7C17F;
  border: 4px solid #F0E9CF;

@media screen and (max-width: 1200px) and (min-width: 1000px) {
  transform:

skewY(10deg

);
}

@media screen and (max-width: 600px) {
  transform:

skewY(11deg

);
}

}
.trapezoid {
  width: 100%;
  height: 300px;
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
  border-top: 0px;
  transform-origin: top;
  box-sizing: border-box;
  z-index: 0;
}

.trapezoid::before {
  border-radius: 0px 0 0 50px;
  border-right: 0;
  left: 0;
  transform: skew(8deg);
}

.trapezoid::after {
  border-radius: 0 50px 50px 0;
  border-left: 0;
  right: 0;
  transform: skew(-8deg);
}

</style>
