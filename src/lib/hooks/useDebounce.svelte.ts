export const useDebounce = <T>(_value: () => T, delay = 500) => {
	let debouncedValueState = $state(_value());

	$effect(() => {
		const newValue = _value();

		const timeoutId = setTimeout(() => {
			debouncedValueState = newValue;
		}, delay);

		return () => clearTimeout(timeoutId);
	});

	return () => debouncedValueState;
};
