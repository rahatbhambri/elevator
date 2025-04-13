<template>
  <div class="container">
    <h2>Elevator Simulation</h2>

    <button v-if="!started" @click="startElevator" class="start-btn">
      Start Elevator
    </button>

    <div v-else class="elevator-shaft">
      <!-- Ropes/Wires -->
      <div class="rope left-rope"></div>
      <div class="rope right-rope"></div>

      <div
        v-for="floor in [...Array(totalFloors).keys()].reverse()"
        :key="floor"
        class="floor"
      >
        <div class="floor-label">
          <!-- Display correct floor numbering, where 0 is at the bottom and 4 is at the top -->
          Floor {{ totalFloors - 1 - floor }}
          <div>
            <button
              v-if="totalFloors - 1 - floor === totalFloors - 1"
              @click="pressButton('down', totalFloors - 1 - floor)"
              :class="{ 'btn-pressed': buttonStates[totalFloors - 1 - floor].down }"
              class="btn"
            >
              ↓
            </button>

            <button
              v-if="totalFloors - 1 - floor === 0"
              @click="pressButton('up', totalFloors - 1 - floor)"
              :class="{ 'btn-pressed': buttonStates[totalFloors - 1 - floor].up }"
              class="btn"
            >
              ↑
            </button>

            <button
              v-if="
                totalFloors - 1 - floor > 0 && totalFloors - 1 - floor < totalFloors - 1
              "
              @click="pressButton('up', totalFloors - 1 - floor)"
              :class="{ 'btn-pressed': buttonStates[totalFloors - 1 - floor].up }"
              class="btn"
            >
              ↑
            </button>
            <button
              v-if="
                totalFloors - 1 - floor > 0 && totalFloors - 1 - floor < totalFloors - 1
              "
              @click="pressButton('down', totalFloors - 1 - floor)"
              :class="{ 'btn-pressed': buttonStates[totalFloors - 1 - floor].down }"
              class="btn"
            >
              ↓
            </button>
          </div>
        </div>
      </div>

      <!-- Elevator -->
      <div
        class="elevator"
        @click="toggleBox"
        :style="{ bottom: elevatorPosition + 'px' }"
      >
        <!-- Left Door -->
        <div class="left-door" :class="{ 'door-open': doorsOpen }"></div>

        <!-- Right Door -->
        <div class="right-door" :class="{ 'door-open': doorsOpen }"></div>

        <!-- Box with floor buttons -->
        <div v-if="isBoxVisible" class="floor-box">
          <button
            v-for="floor in [...Array(totalFloors).keys()]"
            :key="floor"
            @click="goToFloor(floor)"
            class="floor-btn"
          >
            Floor {{ floor }}
          </button>
        </div>
      </div>
    </div>

    <!-- <audio ref="elevatorSound" :src="elevatorSoundFile" preload="auto" /> -->
  </div>
</template>

<script setup>
import { ref } from "vue";
// import elevatorSoundFile from "@/assets/el.mp3";

const totalFloors = 5;
// const elevatorSound = ref(new Audio(elevatorSoundFile));
// const floorHeight = 100
const currentFloor = ref(0);
const elevatorPosition = ref(0);
const started = ref(false);

const floorRequests = ref([]);
const buttonStates = ref(
  Array.from({ length: totalFloors }, () => ({ up: false, down: false }))
);
const doorsOpen = ref(false); // Tracks whether the doors are open
const isBoxVisible = ref(false); // Tracks whether the box is visible

function pressButton(dir, floor) {
  if (!floorRequests.value.includes(floor)) {
    floorRequests.value.push(floor);
    floorRequests.value.sort((a, b) => a - b); // Keep sorted
  }

  // Mark the button as pressed
  if (dir === "up") {
    buttonStates.value[floor].up = true;
  } else if (dir === "down") {
    buttonStates.value[floor].down = true;
  }
}

function toggleBox() {
  isBoxVisible.value = !isBoxVisible.value; // Correctly toggle the visibility of the box
}
function goToFloor(floor) {
  // Close the doors before moving
  doorsOpen.value = false;

  // Update the current floor and elevator position
  currentFloor.value = floor;
  elevatorPosition.value = floor * 100; // Update the elevator's position (assuming each floor is 100px apart)

  // // Play elevator sound if available
  // if (elevatorSound.value) {
  //   elevatorSound.value.currentTime = 0;
  //   elevatorSound.value.play().catch((err) => console.warn("Autoplay blocked:", err));
  // }

  // Open the doors for 1 second after reaching the floor
  setTimeout(() => {
    doorsOpen.value = true;
    setTimeout(() => {
      doorsOpen.value = false; // Close the doors after 1 second
    }, 1000);
  }, 1000); // Delay to simulate the elevator reaching the floor

  // Reset button states for the current floor
  buttonStates.value[floor].up = false;
  buttonStates.value[floor].down = false;

  // Remove the floor from the request queue
  floorRequests.value = floorRequests.value.filter((f) => f !== floor);
}

function getCurrentDirection() {
  // Determine the current direction based on the next request in the queue
  if (floorRequests.value.length === 0) return 0; // No direction
  const nextFloor = floorRequests.value[0];
  return nextFloor > currentFloor.value ? 1 : -1; // 1 for up, -1 for down
}

function simulateElevator() {
  setInterval(() => {
    if (floorRequests.value.length === 0) return;

    const dir = getCurrentDirection(); // Get the current direction of the elevator
    const currFloor = currentFloor.value;

    if (dir === 1) {
      // Moving up
      for (let i = currFloor + 1; i < totalFloors; i++) {
        if (floorRequests.value.includes(i)) {
          goToFloor(i); // Move to the next requested floor
          floorRequests.value = floorRequests.value.filter((f) => f !== i); // Remove served floor
          break; // Stop after serving one floor
        }
      }
    } else if (dir === -1) {
      // Moving down
      for (let i = currFloor - 1; i >= 0; i--) {
        if (floorRequests.value.includes(i)) {
          goToFloor(i); // Move to the next requested floor
          floorRequests.value = floorRequests.value.filter((f) => f !== i); // Remove served floor
          break; // Stop after serving one floor
        }
      }
    }
  }, 3000); // Keep the interval for animation timing
}
function startElevator() {
  started.value = true;
  simulateElevator();
}
</script>

<style scoped>
.rope {
  position: absolute;
  top: 0;
  width: 4px; /* Thickness of the rope */
  height: 100%; /* Extend to the roof */
  background-color: #555; /* Rope color */
  z-index: 1; /* Place behind the elevator */
}

.left-rope {
  left: 45%; /* Position the left rope slightly to the left of the elevator */
}

.right-rope {
  left: 55%; /* Position the right rope slightly to the right of the elevator */
}

.container {
  text-align: center;
  padding: 20px;
}

.start-btn {
  padding: 10px 20px;
  font-size: 16px;
  margin: 20px auto;
  cursor: pointer;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 6px;
}
.btn.btn-pressed {
  background-color: grey !important; /* Ensure it overrides other styles */
  cursor: not-allowed;
}

.elevator-shaft {
  position: relative;
  width: 300px; /* Increase the width to make it broader */
  height: 500px;
  border: 3px solid #444;
  margin: 30px auto;
  background-color: #f5f5f5;
  display: flex;
  flex-direction: column-reverse;
}

.floor {
  height: 100px;
  border-top: 1px solid #ccc;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  padding: 0 10px;
}

.floor-label {
  font-size: 14px;
  color: #333;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.btn {
  background-color: #2196f3;
  color: white;
  border: none;
  padding: 2px 6px;
  border-radius: 4px;
  margin-left: 4px;
  cursor: pointer;
}

.elevator {
  position: absolute;
  left: 50%; /* Position the elevator at the horizontal center of the parent */
  transform: translateX(-50%); /* Adjust to truly center it */
  width: 40%;
  height: 80px;
  background-color: #4caf50;
  border-radius: 8px;
  transition: bottom 1s ease-in-out;
  z-index: 10;
}

.floor-box {
  position: absolute;
  bottom: 100%; /* Position the box above the elevator */
  left: 50%;
  transform: translateX(-50%);
  background-color: white;
  border: 2px solid #ccc;
  border-radius: 8px;
  padding: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  z-index: 20;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.floor-btn {
  background-color: #2196f3;
  color: white;
  border: none;
  padding: 5px 10px;
  margin: 5px 0;
  border-radius: 4px;
  cursor: pointer;
  width: 100px;
  text-align: center;
}

.floor-btn:hover {
  background-color: #1976d2;
}
.left-door,
.right-door {
  position: absolute;
  top: 0;
  height: 100%;
  width: 50%; /* Each door takes half the width of the elevator */
  background-color: #333;
  transition: transform 1s ease-in-out; /* Smooth door animation */
  z-index: 5;
}

.left-door {
  left: 0;
  transform: translateX(0); /* Closed position */
}

.right-door {
  right: 0;
  transform: translateX(0); /* Closed position */
}

.door-open.left-door {
  transform: translateX(-100%); /* Open position for the left door */
}

.door-open.right-door {
  transform: translateX(100%); /* Open position for the right door */
}
</style>
