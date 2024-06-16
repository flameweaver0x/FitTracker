import React from 'react';
import { render, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import FitnessTracker from './FitnessTracker';

process.env.REACT_APP_API_URL = 'https://fakeapi.com';

describe('FitnessTracker Component Tests', () => {
  test('renders fitness tracker component', () => {
    const { getByText } = render(<FitnessTracker />);
    expect(getByText('Fitness Tracker')).toBeInTheDocument();
  });

  test('input value updates on change', () => {
    const { getByPlaceholderText } = render(<FitnessTracker />);
    const inputField = getByPlaceholderText('Enter Steps');
    fireEvent.change(inputField, { target: { value: '1000' } });
    expect(inputField.value).toBe('1000');
  });

  test('displays real-time updates', async () => {
    const { getByTestId } = render(<FitnessTracker />);
    const realTimeDataDisplay = getByTestId('real-time-data');
    await waitFor(() => expect(realTimeDataDisplay).toHaveText>Content('Updated Steps: 5000'));
  });

  test('click event triggers state update', () => {
    const { getByText, getByTestId } = render(<FitnessTracker />);
    fireEvent.click(getByText('Update Steps'));
    const stepsDisplay = getByTestId('steps-display');
    expect(stepsDisplay).toHaveTextContent('Steps: 1000');
  });
});