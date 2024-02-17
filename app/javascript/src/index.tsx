import React from 'react';
import { createRoot } from 'react-dom/client';
import TestComponent from './TestComponent';

// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
const root = createRoot(document.getElementById("app")!);

root.render(<TestComponent />);
