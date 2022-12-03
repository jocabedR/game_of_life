<template>
  <h1>Game of life</h1>

  <button @click="setDimensions(10)">10x10</button>
  <button @click="setDimensions(20)">20x20</button>
  <button @click="setDimensions(50)">50x50</button>
  <button @click="setDimensions(100)">100x100</button>

  <form class="options" @submit.prevent="setGrid">
    <br>Probability: <input type="number" min="0" max="1" step="0.1" v-model="probability"/>
    <br><button>Play!</button>
  </form>
  
  <Grid :grid="grid"/>
  <!-- <HelloWorldVue/> -->

</template>

<script>
import axios from 'axios'
import Grid from './components/Grid.vue'
//import HelloWorldVue from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    Grid
  },

  data() {
    return{
      size: 0,
      probability: 0.0,
      grid: [],
    }
  },

  methods:{
    setDimensions(size){
      this.size = size
    },
    setGrid() {
      this.grid = []
      for(let i=0; i<this.size; i++){
        this.grid[i] = []
        for (let j = 0; j < this.size; j++) {
          if(this.randomNumber() > (this.probability)*10){
            this.grid[i][j] = false 
          } 
          else{
            this.grid[i][j] = true
          } 
        }
      }
      this.postreq()
    },
    randomNumber() {
      return Math.floor(Math.random() * (10 - 1 + 1)) + 1
    },

    postreq: function() {
      var data = {"grid": this.grid}

      console.log(this.data)
      
      axios({ method: "POST", url: "http://127.0.0.1:3000/game", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
          //console.log(result.data)

          this.grid = result.data

          console.log(this.grid)

        }).catch( error => {
            console.error(error);
      });
    } 
  }
}
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 3rem;
  }

  .number-input, button{
    margin: 1rem 0.25rem;
  }
</style>