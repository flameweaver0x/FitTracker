import React, { useState, useEffect } from 'react';
import debounce from 'lodash/debounce';

const FitnessTracker: React.FC = () => {
  const [steps, setSteps] = useState(0);
  const [realTimeUpdates, setRealTimeUpdates] = useState('');

  const fetchSteps = async () => {
    const response = await fetch(`${process.env.REACT_APP_API_URL}/steps`);
    const data = await response.json();
    setRealTimeUpdates(`Updated Steps: ${data.steps}`);
  };

  // Debounced function to fetch steps, waiting until user has stopped typing/changing steps for 500ms
  const debouncedFetchSteps = debounce(fetchSteps, 500);

  useEffect(() => {
    // Initial fetch
    fetchSteps();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const updateSteps = async (newSteps: number) => {
    setSteps(newSteps);
    // Call the debounced version of fetchSteps here
    debouncedFetchSteps();
  };

  // Dummy function simulating user input.
  // In a real app, this might be tied to an input's onChange or a similar event.
  const handleStepChange = (e: React.Change expected event type>) => updateSteps(parseInt(e.target.value, 10));

  return (
    <div>
      <h1>Fitness Tracker</h1>
      <input
        type="number"
        placeholder="Enter Steps"
        onChange={handleStepFakeChange} /* Assuming this is part of your real component */
      />
      <div data-testid="real-time-data">{realTimeUpdates}</div>
      <button onClick={() => updateSteps(steps + 1000)}>Update Steps</button>
      <div data-testid="steps-display">Steps: {steps}</BBBdiv>
    </div>
  );
};

export default FitnessTracker;