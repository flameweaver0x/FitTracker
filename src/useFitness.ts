import { useState, useEffect } from 'react';
import axios from 'axios';

interface Workout {
  id: string;
  date: string;
  exercises: Exercise[];
}

interface Exercise {
  name: string;
  sets: number;
  reps: number;
  weight: number;
}

const useFitnessData = () => {
  const [workoutHistory, setWorkoutHistory] = useState<Workout[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const baseURL = process.env.REACT_APP_BACKEND_URL;

  const fetchWorkoutHistory = async (): Promise<void> => {
    setLoading(true);
    try {
      const { data } = await axios.get(`${baseURL}/workouts`);
      setWorkoutHistory(data);
      setLoading(false);
    } catch (err) {
      setError('Failed to fetch workout history');
      setLoading(false);
    }
  };

  const updateExerciseLog = async (workoutId: string, exercise: Exercise): Promise<void> => {
    try {
      await axios.post(`${baseURL}/workouts/${workoutId}/exercises`, exercise);
      fetchWorkoutHistory();
    } catch (err) {
      setError('Failed to update exercise log');
    }
  };

  const fetchWorkoutRecommendations = async (): Promise<Workout[] | undefined> => {
    try {
      const { data } = await axios.get(`${baseURL}/recommendations`);
      return data;
    } catch (err) {
      setError('Failed to fetch workout recommendations');
      return undefined;
    }
  };

  useEffect(() => {
    fetchWorkoutHistory();
  }, []);

  return {
    workoutHistory,
    fetchWorkoutHistory,
    updateExerciseLog,
    fetchWorkoutRecommendations,
    loading,
    error,
  };
};

export default useFitnessData;