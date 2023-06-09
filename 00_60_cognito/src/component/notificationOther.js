import * as React from 'react';
import Box from '@mui/material/Box';
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import Button from '@mui/material/Button';

export default function ButtonsTransition() {
  const [checked, setChecked] = React.useState(false);

  function handleClick() {
    setChecked(true);
  }

  return (
    <Box>
    <Box sx={{ '& > button': { m: 1 } }}>
      <Button
        size="Large"
        onClick={handleClick}
        checked={checked}
        variant="text"
        disabled={checked ? false : true}
      >
        <label>その他（Line、Teams)</label>
      </Button>
    </Box>

      <FormControlLabel
        sx={{
          display: 'block',
          marginLeft: '5px',
        }}
        control={
          <Switch
            checked={checked}
            onChange={() => setChecked(!checked)}
            name=""
            color="primary"
          />
        }
        label={`${checked ? "オン" : "オフ"}`}
      />
    </Box>
  );
}