<template>
  <div v-if="showModal" class="modal-overlay" @click="closeModal">
      <div class="modal" @click.stop>
        <h2>Добавление тура</h2>

        <form @submit.prevent="submitForm">
          <div style="display:flex; flex-direction: column;">
            <input type="text" placeholder="Название тура" v-model="tour.title" required />
            <input type="text" placeholder="Описание" v-model="tour.description" required />
            <input type="number" placeholder="Цена в рублях" v-model="tour.price" required />
            <!-- <input type="file" @change="onFileChange" accept="image/*" required /> -->
          </div>
          <button type="submit">Отправить</button>
          <button type="button" @click="closeModal">Закрыть</button>
        </form>
      </div>
  </div>
    
  <div class="wrapper">
    <h1 class="RegAuth">Туры</h1>
  </div>

  <button class="addTour" @click="showModal = true">Добавить тур</button>
  <div class="tourBlock">
    <div class="tourCards" v-for="tour in Tours" :key="tour.id">
      <TourItem :placeTour="tour.title" :price="tour.price" :description="tour.description"/>
    </div>
  </div>
</template>

<script>
import TourItem from './components/TourItem.vue'
import axios from 'axios';

export default {
  components: {
    TourItem 
  },
  data() {
    return {
      Tours: [],
      error: null, 
      showModal: false,
      tour : {title: '', description: '', price: ''},
    };
  },
  mounted() {
    axios.get('http://localhost:8084/tours')
      .then(response => {
        this.Tours = response.data;
      })
      .catch(err => {
        this.error = 'Ошибка при загрузке задач: ' + err.message; 
      });
  },
  methods: {
        closeModal() {
            this.showModal = false
            this.tour = { title: '', description: '', price: null }
        },
        async submitForm() {
          try {
            const response = await axios.post('http://localhost:8084/tours', this.tour);
            this.Tours.push(response.data); 
            this.closeModal(); 
          } catch (error) {
            console.error('Error adding tour:', error);
          }
        },
      }
};
</script>

<style>
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.7); /* Затемнение фона */
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal {
    background-color: white;
    padding: 20px;
    border-radius: 17px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    height: 60%;
    width: 50%; /* Ширина модального окна */
}
.modal h2{
  color: black;
  font-size: 30px;
  font-weight: 700;
  margin-bottom: 5%;
}
form input{
  padding: 10px;
  width: 70%;
  border-radius: 17px;
  margin-bottom: 2%;
}
form button{
  cursor: pointer;
  border-radius: 17px;
  background-color: white;
  padding: 10px;
  margin-right: 2%;
}
.wrapper{
  width: 100%;
  display: flex;
  justify-content: center;
}
.addTour{
  cursor: pointer;
  padding: 13px;
  background-color: #fff;
  font-size: 15px;
  border-radius: 18px;
  margin-top: 10px;
  border: 0;
  margin-bottom: 2%;
}
.RegAuth{
  color: black;
  font-size: 30px;
  font-weight: 700;
}
.tourBlock{
  display: grid;
  grid-template-columns: repeat(4, 1fr); 
  gap: 20px;
}
.tourCards{ 
  margin-bottom: 20px;
  margin-right: 10%
}
.error {
  color: red;
}

</style>