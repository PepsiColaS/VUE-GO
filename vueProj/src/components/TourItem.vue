<template>
  <div>
    <div class="tourCard">
      <div style="display: flex; justify-content: center;">
        <img src="../../public/MoscowImg.avif" alt="tourImg">
      </div>
      
      <div class="info">
        <h4>{{ placeTour }}</h4>
        <p>{{ price }} руб</p>
      </div>

      <button class="btnBy" @click="showDetails">Подробнее</button>
    </div>

    <!-- Модальное окно с подробностями о туре -->
    <div v-if="showDetailsModal" class="modal-overlay" @click="closeDetailsModal">
      <div class="modal" @click.stop>
        <h2>Подробности о туре</h2>
        <p><strong>Название:</strong> {{ placeTour }}</p>
        <p><strong>Описание:</strong> {{ description }}</p>
        <p><strong>Цена:</strong> {{ price }} руб</p>
        <button @click="closeDetailsModal">Закрыть</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    placeTour: {
      type: String,
      required: true
    },
    price: {
      type: String,
      required: true
    },
    description: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      showDetailsModal: false // Состояние модального окна для подробностей о туре
    };
  },
  methods: {
    showDetails() {
      this.showDetailsModal = true; // Открыть модальное окно с подробностями о туре
    },
    closeDetailsModal() {
      this.showDetailsModal = false; // Закрыть модальное окно с подробностями о туре
    }
  }
};
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
}

.info {
  margin-top: 10px;
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

/* Стили для модального окна */
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

.modal {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 300px;
}
</style>