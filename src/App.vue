<template>
  <h1>Game of life</h1>
  <form class="options" @submit.prevent="setGrid">
    Width: <input class="number-input" type="number"  v-model="width" /> 
    <br>Height: <input class="number-input" type="number" v-model="height"/> 
    <br>Probability: <input type="number" min="0" max="1" step="0.1" v-model="probability"/>
    <br><button>Play!</button>
  </form>
  
  <table>
    <tr v-for="(row, rowKey) in cells" :key="rowKey">
      <td v-for="(col, colKey) in row" :key="colKey" 
      :class="col.alive ? 'alive' : 'dead'" class="cell"
      ></td>
    </tr>
  </table>
</template>

<script>
let id = 0;
export default {
  name: 'App',
  components: {
  },

  data() {
    return{
      width: 0,
      height: 0, 
      probability: 0.0,
      cells: []
    }
  },

  methods:{
    setGrid() {
      this.cells = []
      id = 0
      for(let i=0; i<this.height; i++){
        this.cells[i] = []
        for (let j = 0; j < this.width; j++) {
          if(this.randomNumber() > (this.probability)*10) this.cells[i][j]={ id: id++, alive: false }
          else this.cells[i][j]={ id: id++, alive: true }
        }
      }

    },
    changeState(cell, i,j) {
      this.cells[i][j].alive = !cell.alive
      console.log(cell)
    },
    
    randomNumber() {
      return Math.floor(Math.random() * (10 - 1 + 1)) + 1
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
  margin-top: 3em;
}

.number-input, button{
  margin: 0.25em
}

.cell {
  border-style:solid;
  padding: 0.25em;
}

/* .cell:hover {
  background-color: gray ;
} */

.alive {
  background-color: black !important;
}

.dead {
  background-color: white ;
}

table {
  margin: auto;
}

</style>