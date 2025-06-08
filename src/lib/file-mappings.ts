export const mimeTypeTextMappings: Map<string, string> = new Map([
	['application/pdf', 'PDF'],
	['application/vnd.siemens.tiaportal.project.archive', 'TIA Portal Project Archive'],
	['application/zip', 'ZIP Archive'],
	['application/x-zip-compressed', 'ZIP Archive']
]);

export type MimeTypeIconMapping =
	| 'image'
	| 'video'
	| 'audio'
	| 'document'
	| 'archive'
	| 'code'
	| 'unknown';
export const mimeTypeIconMappings: Map<string, MimeTypeIconMapping> = new Map([
	['application/pdf', 'document'],
	['application/vnd.siemens.tiaportal.project.archive', 'unknown'],
	['application/zip', 'archive'],
	['application/x-zip-compressed', 'archive']
]);

export const resolveTextMapping = (mimeType: string): string => {
	if (mimeTypeTextMappings.has(mimeType)) {
		return mimeTypeTextMappings.get(mimeType)!;
	}

	if (mimeType.startsWith('image/')) {
		return 'Image';
	}

	if (mimeType.startsWith('video/')) {
		return 'Video';
	}

	if (mimeType.startsWith('audio/')) {
		return 'Audio';
	}

	if (mimeType.startsWith('text/')) {
		return 'Text File';
	}

	return 'File';
};

export const resolveIconMapping = (mimeType: string): MimeTypeIconMapping => {
	if (mimeTypeIconMappings.has(mimeType)) {
		return mimeTypeIconMappings.get(mimeType)!;
	}

	if (mimeType.startsWith('image/')) {
		return 'image';
	}

	if (mimeType.startsWith('video/')) {
		return 'video';
	}

	if (mimeType.startsWith('audio/')) {
		return 'audio';
	}

	if (mimeType.startsWith('text/')) {
		return 'code';
	}

	return 'unknown';
};
