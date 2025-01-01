<template>
  <div>
    <div class="tourCard">
      <div style="display: flex; justify-content: center;">
      </div>
      
      <div class="info">
        <h4>{{ placeTour }}</h4>
        <p>{{ price }} руб</p>
      </div>

      <div class="wrap">
        <button class="btnBy" @click="showDetails">Подробнее</button>
        <button class="btnBy" @click="remove">Удалить</button>
      </div>
      
    </div>

    <!-- Вылетающая формочка доп информации о туре -->
    <div v-if="showDetailsModal" class="modal-overlay" @click="closeDetailsModal">
      <div class="modal" @click.stop>
        <h2 style="margin-bottom: 20px;">Подробности о туре</h2>
        <p><strong>Название:</strong> <input class="inputDescriprion" type="text" v-model="tour.title" :placeholder="placeTour"></p>
        <p><strong>Описание:</strong> <input class="inputDescriprion" type="text" v-model="tour.description" :placeholder="description"></p>
        <p><strong>Цена:</strong> <input class="inputDescriprion" type="text" v-model="tour.price" :placeholder="`${price} руб`"></p>
        <button class="btnCloseSave" @click="closeDetailsModal">Закрыть</button>
        <button class="btnCloseSave" @click="saveEditDetails">Сохранить</button>
      </div>
    </div>
    <!-- ---- -->

  </div>
</template>

<script>

export default {
  props: {
    placeTour: { type: String, required: true },
    price: { type: String, required: true },
    description: { type: String, required: true },
    tourId: { type: Number, required: true },
    deleteTour: Function,
    updateTour: Function,
  },
  data() {
    return {
      showDetailsModal: false,
      tour : {id: null, title: '', description: '', price: ''}
    }
  },
  created() {
    this.tour.id = this.tourId;
    this.tour.title = this.placeTour;
    this.tour.description = this.description;
    this.tour.price = this.price;
  },
  methods: {
    showDetails() {
      this.showDetailsModal = true
    },
    closeDetailsModal() {
      this.showDetailsModal = false
    },
    saveEditDetails(){
      this.updateTour(this.tour)
      this.closeDetailsModal()
    },
    remove(){
      this.deleteTour(this.tourId)
    }
  }
}
</script>

<style>
.tourCard {
  background-color: #fff;
  padding: 20px;
  border-radius: 18px; 
  width: 100%;
}

.tourCard img {
  width: 100%;
  min-height: 170px;
  max-height: 200px;
  border-radius: 18px;
}

h4 {
  color: black;
  font-weight: bold;
}

p {
  color: black;
  margin-bottom: 8px;
}

.info {
  margin-top: 10px;
}

.wrap{
  display: flex;
  justify-content: space-between;
}

.btnBy {
  cursor: pointer;
  padding: 10px;
  background-color: #edecec;
  font-size: 12px;
  border-radius: 18px;
  margin-top: 10px;
  border: 0;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

strong{
  margin-right: 5px;
}

.inputDescriprion{
  padding: 8px;
  border-radius: 17px;
  border: 1px solid black;
}

.inputDescriprion::placeholder{
  color: black;
}

.modal {
  background: white;
  padding: 10px;
  border-radius: 8px;
  width: 300px;
}

.btnCloseSave{
  font-size: small;
  cursor: pointer;
  padding: 12px;
  margin-top: 2%;
  border: 0;
  border-radius: 17px;
  margin-right: 8px;
}
</style>