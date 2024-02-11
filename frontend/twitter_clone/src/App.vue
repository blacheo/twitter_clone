<script setup>
import SearchBar from './components/SearchBar.vue';
import Tweet from './components/Tweet.vue';
import axios from 'axios';
import {ref, onMounted as onMounted} from 'vue';

const tweets = ref(null)


onMounted(() => {
axios
.get('http://localhost:5000/tweets')
.then(response => {tweets.value = response.data})
.catch(error => {
  console.log(error)
  tweets.errored = true
}
)


})


</script>

<template>
  <header>

    <div class="wrapper">
      
    </div>
  </header>

  <main>
    <h1>Twitter Clone</h1>
    <SearchBar/>
    
      <template v-for="tweet in tweets">
        <Tweet v-bind:user="tweet.user_id" v-bind:content="tweet.content"/>
      </template>
    
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
