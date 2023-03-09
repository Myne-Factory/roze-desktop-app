export namespace main {
	
	export class Image {
	    name: string;
	    drive: string;
	    path: string;
	    fullPath: string;
	    copied: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.drive = source["drive"];
	        this.path = source["path"];
	        this.fullPath = source["fullPath"];
	        this.copied = source["copied"];
	    }
	}

}

