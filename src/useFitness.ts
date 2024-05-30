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

const useWorkoutData = () => {
  const [workoutLogs, setWorkoutLogs] = useState<Workout[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [fetchError, setFetchError] = useState<string | null>(null);

  const apiBaseUrl = process.env.REACT_APP_BACKEND_URL;

  const fetchWorkoutLogs = async (): Promise<void> => {
    setIsLoading(true);
    try {
      const { data } = await axios.get(`${apiBaseUrl}/workouts`);
      setWorkoutLogs(data);
      setIsLoading(false);
    } catch (error) {
      setFetchError('Failed to fetch workout logs');
      setIsLoading(false);
    }
  };

  const addExerciseToWorkout = async (workoutId: string, newExercise: Exercise): Promise<void> => {
    try {
      await axios.post(`${apiBaseUrl}/workouts/${workoutId}/exercises`, newExercise);
      fetchWorkoutLogs(); // Refresh the workout logs to include the newly added exercise
    } catch (error) {
      setFetchError('Failed to add exercise to workout log');
    }
  };

  const fetchRecommendedWorkouts = async (): Promise<Workout[] | undefined> => {
    try {
      const { data } = await axios.get(`${apiBaseUrl}/recommendations`);
      return data;
    } catch (error) {
      setFetchError('Failed to fetch workout recommendations');
      return undefined;
    }
  };

  useEffect(() => {
    fetchWorkoutLogs();
  }, []);

  return {
    workoutLogs,
    fetchWorkoutLogs,
    addExerciseToWorkout,
    fetchRecommendedWorkouts,
    isLoading,
    fetchError,
  };
};

export default useWorkoutData;