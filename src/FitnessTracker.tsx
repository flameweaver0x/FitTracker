import React, { useState, useEffect } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
import DatePicker from 'react-datepicker';
import "react-datepicker/dist/react-datepicker.css";

interface Workout {
  date: Date;
  duration: number; 
  intensity: number; 
}

const FitnessTracker: React.FC = () => {
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [newWorkout, setNewWorkout] = useState<Workout>({ date: new Date(), duration: 0, intensity: 0 });
  const [fitnessGoal, setFitnessGoal] = useState<number>(0);

  const handleNewWorkoutChange = (e: React.ChangeEvent<HTMLInputElement> | React.ChangeEvent<HTMLSelectElement>, field: string) => {
    setNewWorkout({ ...newWorkout, [field]: field === "date" ? new Date(e.target.value) : parseInt(e.target.value) });
  };

  const logWorkout = () => {
    if (newWorkout.duration > 0 && newWorkout.intensity > 0) {
      setWorkouts([...workouts, newWorkout]);
      setNewWorkout({ date: new Date(), duration: 0, intensity: 0 }); 
    }
  };

  const totalDuration = workouts.reduce((acc, workout) => acc + workout.duration, 0);

  const isGoalAchieved = totalDuration >= fitnessGoal;

  return (
    <div>
      <h2>Fitness Tracker</h2>
      <div>
        <label>
          Select Date:
          <DatePicker selected={newWorkout.date} onChange={(date: Date) => setNewWorkout({ ...newWorkout, date })} />
        </label>
        <label>
          Duration (in minutes):
          <input type="number" value={newWorkout.duration} onChange={(e) => handleNewWorkoutChange(e, 'duration')} />
        </label>
        <label>
          Intensity (1-10):
          <input type="number" min="1" max="10" value={newWorkout.intensity} onChange={(e) => handleNewWorkoutChange(e, 'intensity')} />
        </label>
        <button onClick={logWorkout}>Log Workout</button>
      </div>
      <div>
        <label>
          Set your fitness goal (in minutes):
          <input type="number" value={fitnessGoal} onChange={(e) => setFitnessGoal(parseInt(e.target.value))} />
        </label>
        {isGoalAchieved && <div>Congratulations! You've achieved your goal!</div>}
      </div>
      <LineChart width={500} height={300} data={workouts}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="date" />
        <YAxis />
        <Tooltip />
        <Legend />
        <Line type="monotone" dataKey="duration" stroke="#8884d8" />
        <Line type="monotone" dataKey="intensity" stroke="#82ca9d" />
      </LineChart>
    </div>
  );
};

export default FitnessTracker;