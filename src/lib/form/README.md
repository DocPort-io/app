# Custom Form System

A type-safe, reactive form system built with Svelte 5 runes and Zod validation.

## Features

- 🔒 **Type-safe**: Full TypeScript support with schema inference
- ⚡ **Reactive**: Built with Svelte 5 runes for optimal performance
- 🛡️ **Validation**: Powered by Zod schemas with field-level and form-level validation
- 🎨 **Composable**: Modular component architecture
- 🚀 **Developer Experience**: Intuitive API with great autocomplete support

## Quick Start

### 1. Define your schema

```typescript
import * as z from 'zod';

const userSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  email: z.email('Invalid email'),
  age: z.number().min(18, 'Must be at least 18')
});
```

### 2. Create a form

```svelte
<script lang="ts">
  import { createForm } from '$lib/form';
  import Field from '$lib/form/field.svelte';
  import FieldLabel from '$lib/form/field-label.svelte';
  import FieldErrors from '$lib/form/field-errors.svelte';
  import { Input } from '$lib/components/ui/input';

  const form = createForm({
    schema: userSchema,
    defaultValues: {
      name: '',
      email: ''
    },
    onSubmit: async ({ data, state }) => {
      console.log('Form data:', data);
      // Handle submission
    }
  });
</script>

<form {...form.props}>
  <Field {form} name="name">
    {#snippet children({ props, state })}
      <FieldLabel>Name</FieldLabel>
      <Input {...props} bind:value={state.value} />
      <FieldErrors />
    {/snippet}
  </Field>

  <Field {form} name="email">
    {#snippet children({ props, state })}
      <FieldLabel>Email</FieldLabel>
      <Input {...props} bind:value={state.value} type="email" />
      <FieldErrors />
    {/snippet}
  </Field>

  <button type="submit" disabled={!form.state.isSubmittable}>
    {form.state.isSubmitting ? 'Submitting...' : 'Submit'}
  </button>
</form>
```

## API Reference

### `createForm(options)`

Creates a new form instance.

#### Options

- `schema`: Zod schema defining the form structure
- `defaultValues?`: Initial values for form fields
- `onSubmit?`: Function called when form is submitted
- `validateOnChange?`: Enable validation on field change (default: `true`)
- `validateOnBlur?`: Enable validation on field blur (default: `true`)

#### Returns

A form object with:

- `state`: Reactive form state
- `fields`: Field objects for each schema property
- `props`: Props to spread on the form element
- `validate()`: Manually trigger validation
- `reset()`: Reset form to initial state
- `setFieldValue()`: Update a field value
- `getFieldValue()`: Get a field value

### Form State

- `isSubmittable`: Whether the form can be submitted
- `isSubmitting`: Whether the form is currently submitting
- `isValidating`: Whether validation is in progress
- `isValid`: Whether the form is valid
- `errors`: Object containing field errors
- `hasErrors`: Whether the form has any errors

### Components

#### `<Field>`

Wrapper component for form fields.

```svelte
<Field {form} name="fieldName">
  {#snippet children({ props, state })}
    <!-- Field content -->
  {/snippet}
</Field>
```

#### `<FieldLabel>`

Label component that automatically links to the field.

```svelte
<FieldLabel>Field Label</FieldLabel>
```

#### `<FieldErrors>`

Displays field validation errors.

```svelte
<FieldErrors />
```

#### `<FieldDescription>`

Helper text for fields.

```svelte
<FieldDescription>Field description</FieldDescription>
```

## Advanced Usage

### Custom Validation

```typescript
const form = createForm({
  schema: userSchema,
  validateOnChange: false, // Disable real-time validation
  onSubmit: async ({ data, state }) => {
    // Manual validation
    const isValid = await form.validate();
    if (!isValid) return;
    
    // Submit data
  }
});
```

### Complex Schema Examples

```typescript
import * as z from 'zod';

// Nested objects
const addressSchema = z.object({
  street: z.string().min(1, 'Street is required'),
  city: z.string().min(1, 'City is required'),
  zipCode: z.string().regex(/^\d{5}$/, 'Invalid zip code')
});

// Optional and nullable fields
const profileSchema = z.object({
  name: z.string().min(1),
  bio: z.string().optional(),
  avatar: z.url().nullable(),
  age: z.number().int().min(13).max(120)
});

// Arrays and enums
const preferencesSchema = z.object({
  tags: z.array(z.string()).min(1, 'At least one tag required'),
  status: z.enum(['active', 'inactive', 'pending']),
  notifications: z.boolean().default(true)
});

// String formats and validation
const contactSchema = z.object({
  email: z.email(),
  phone: z.string().regex(/^\+?[\d\s-()]+$/, 'Invalid phone format'),
  website: z.url().optional(),
  birthDate: z.string().regex(/^\d{4}-\d{2}-\d{2}$/, 'Use YYYY-MM-DD format')
});
```

### Field-level Control

```typescript
// Set field value programmatically
form.setFieldValue('name', 'John Doe');

// Get field value
const name = form.getFieldValue('name');

// Reset form
form.reset();
```

### Real-time Validation Options

```typescript
import * as z from 'zod';

const form = createForm({
  schema: z.object({
    username: z.string().min(3, 'Username too short'),
    password: z.string().min(8, 'Password must be at least 8 characters')
  }),
  validateOnChange: true,  // Validate as user types
  validateOnBlur: true,    // Validate when field loses focus
  validateOnInput: true    // Validate on every input event
});
```

### Custom Field Components

```svelte
<Field {form} name="status">
  {#snippet children({ props, state })}
    <FieldLabel>Status</FieldLabel>
    <Select {...props} bind:value={state.value}>
      <!-- Select options -->
    </Select>
    {#if state.hasErrors}
      <FieldErrors />
    {/if}
  {/snippet}
</Field>
```

## Benefits over FormSnap/SuperForms

1. **Simpler API**: Less boilerplate and more intuitive
2. **Better TypeScript**: Native type inference without complex generics
3. **Svelte 5 Optimized**: Built specifically for Svelte 5 runes
4. **Lighter Weight**: Fewer dependencies and smaller bundle size
5. **More Control**: Direct access to field state and validation

## Zod v4 Integration

This form system is built to take full advantage of Zod v4's enhanced features:

### Modern Import Syntax
```typescript
import * as z from 'zod';  // Recommended v4 import pattern
```

### Enhanced Validation APIs
- **String formats**: `z.email()`, `z.url()`, `z.uuid()`, `z.ipv4()`, etc.
- **Template literals**: `z.templateLiteral(['hello, ', z.string(), '!'])`
- **String transforms**: `z.string().trim().toLowerCase()`
- **Improved enums**: Better TypeScript enum support with `z.enum()`
- **Branded types**: Type-safe nominal typing with `.brand<'Type'>()`

### Performance Improvements
- **Discriminated unions**: Faster parsing with `z.discriminatedUnion()`
- **Lazy evaluation**: Better handling of recursive schemas
- **Optimized error handling**: More precise error messages and paths

### New Validation Features
```typescript
// String boolean validation
const configSchema = z.object({
  debug: z.stringbool(),  // "true"/"false" → boolean
  enabled: z.stringbool({ truthy: ['yes', '1'], falsy: ['no', '0'] })
});

// ISO date/time validation
const scheduleSchema = z.object({
  date: z.iso.date(),     // Uses z.iso.date() for YYYY-MM-DD
  time: z.iso.time()      // Uses z.iso.time() for HH:MM
});

// File validation
const uploadSchema = z.object({
  avatar: z.instanceof(File)
    .refine(file => file.size <= 5_000_000, 'File too large')
    .refine(file => file.type.startsWith('image/'), 'Must be an image')
});
```
